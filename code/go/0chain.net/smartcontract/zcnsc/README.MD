![Diagram](https://github.com/Dmdv/0chain/blob/zcnsc_latest/code/go/0chain.net/smartcontract/zcnsc/docs/ZCNSC.png?raw=true)


## Converting WZCN to ZCN

1. The client owning WZCN writes a transaction to the WZCN mint to burn  
   tokens; we refer to this transaction as the WZCN burn transaction. This  
   transaction contains:  
   • The amount of WZCN burned.  
   • The client’s ID on the 0chain blockchain.  
   • A sequential nonce. (This nonce is distinct from the nonce used in Ethereum’s protocol.)  
2. The Ethereum network accepts the transaction and includes it in the blockchain.  
   Note that it will only be accepted if the nonce is one greater than the previous nonce.  
   The authorizers monitor the Ethereum blockchain for WZCN burn transactions.  
3. Each authorizer verifies that the transaction has been accepted on the  
   Ethereum blockchain. If the request is valid, the authorizer sends the  
   client a proof-of-WZCN-burn ticket. This ticket contains:  
   • The Ethereum transaction ID.  
   • The amount of ZCN to be minted.  
   • The client ID to receive the ZCN.  
   • The nonce provided the client in step 1 above.  
   • The authorizer’s signature.  
4. Once the client has gathered sufficient tickets, they write a ZCN mint  
   transaction containing these tickets. The 0chain miners verify the validity  
   of the tickets, checking the signatures and the nonces. If the ZCN mint  
   transaction is valid, the miners accept the transaction and allocate new  
   ZCN to the client’s account.  

## Converting ZCN to WZCN

The reverse process is similar. A client burns ZCN and presents proof to the  
authorizers, who then authorize the generation of new WZCN tokens.  
The steps for this process are as follows:  

1. The client writes a ZCN burn transaction, destroying ZCN tokens. This  
   transaction includes:  
   • The amount of ZCN to burn.  
   • The Ethereum address to receive the new WZCN.  
   • The nonce value.  
2. The 0Chain miners accept this transaction if the nonce is valid and the  
   client has sufficient ZCN  
3. Each authorizer monitors the 0Chain blockchain for a ZCN burn transac-  
   tion. Once the transaction is accepted, the authorizer sends the client a  
   proof-of-ZCN-burn ticket, containing similar information as the proof-of-  
   WZCN-burn ticket:  
   • The 0Chain transaction ID.  
   • The nonce value.  
   • The amount of WZCN to be minted.  
   • The Ethereum address to receive the WZCN.  
   • The authorizer’s signature.  
4. Once the client has gathered sufficient tickets, they write a transaction to  
   the WZCN mint, including the tickets from the authorizers.  
5. The WZCN mint verifies the validity of the tickets, checking the signatures  
   and the nonce values. If the transaction is valid, the WZCN mint creates  
   new WZCN for the client.  