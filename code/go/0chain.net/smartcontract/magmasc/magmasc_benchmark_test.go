package magmasc

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"0chain.net/smartcontract/magmasc/bench-state-gen/dir"
)

func BenchmarkMagmaSmartContract_Execute(b *testing.B) {
	var (
		nas, nis    = countSessions(dir.SciDir, dir.SciLogDir, dir.DbDir, b)
		sessPostfix = "_" + strconv.Itoa(nas) + "as_" + strconv.Itoa(nis) + "is"

		consumers, providers = countNodes(dir.SciDir, dir.SciLogDir, dir.DbDir, b)
		nodesPostfix         = "_" + strconv.Itoa(consumers) + "c_" + strconv.Itoa(providers) + "p"

		cpSciDir, cpSciLogDir, cpDbDir = b.TempDir(), b.TempDir(), b.TempDir()
	)
	copyDir(dir.SciDir, cpSciDir, b)
	copyDir(dir.DbDir, cpDbDir, b)
	copyDir(dir.SciLogDir, cpSciLogDir, b)

	for _, tCase := range loadBenchExecuteCases() {
		// stress
		name := b.Name() + "_Stress_" + tCase.funcName() + sessPostfix + nodesPostfix
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				tCase.refresh()

				var (
					sc = NewMagmaSmartContract()

					txn, funcName, input = tCase.txn(), tCase.funcName(), tCase.input()

					cpDb, cpSci = tCase.setupDatabaseAndState(cpDbDir, cpSciDir, cpSciLogDir, b)
				)
				sc.db = cpDb

				b.StartTimer()
				_, err := sc.Execute(txn, funcName, input, cpSci)
				b.StopTimer()

				require.NoError(b, err)

				closeSciAndDB(cpSci, cpDb, b)
			}
		})

		name = b.Name() + "_" + tCase.funcName()
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				tCase.refresh()

				var (
					sc = NewMagmaSmartContract()

					txn, funcName, input = tCase.txn(), tCase.funcName(), tCase.input()

					db, sci = tCase.setupDatabaseAndState(b.TempDir(), b.TempDir(), b.TempDir(), b)
				)
				sc.db = db

				b.StartTimer()
				_, err := sc.Execute(txn, funcName, input, sci)
				b.StopTimer()

				require.NoError(b, err)
				closeSciAndDB(sci, db, b)
			}
		})
	}
}
