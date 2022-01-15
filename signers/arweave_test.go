package signers

import (
	"crypto/sha256"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArweave_Sign(t *testing.T) {
	var walletJWK = `{
    "p": "5sbDW0jFyxb2ewKkhpAhB2N8Rv-gglD3zWcgCJIcF0Q2vUlWUd07YKvhc6hLQq0FDScoD_E1CkywqYXFSsXu6oTP_dSKpK8wmUhzYDt3CqD2fr0FmoogQXhirxVuTfDKvX0MnwRpQeGR_M-TODGKoTs5tO89xqhwt5N2NsQi-8M",
    "kty": "RSA",
    "q": "zuZMWm1f5YBHYyPtptnMv3oyiUAFbmy7RNKNx3ck0fDuL_9CcpctKeniSlxRoNWxw808XwSmA1HyOERWacnAmbKhJOztKgK_14CqUMNPec86Nyfq6--h-sfUk6O0rmQm34aGflTIIVhStQQ7MSocvBTzNAeYE403s5pwfx2DLN0",
    "d": "pmvs0W6J0L2rNBUa_dLurWTV2xovvDJrnuDLOIwf63I9-KKkifdcwLLauy2vz7tAs2qSD5L8CXMpuWgymx_oQqu6depfACQhQqrYfJ4x3QfyB_dgxEqNMxGhg6U1W9eOWUOoagjEyA9O21AsgtF3xvXM-3IHBPPDPmNfPibhqcpIPd8L0Q9Vr1T65YrNsIUDbPKgWcwpk6hAzq-H31_Gjy4AteWFF6KINj7TAfMh_IRRNApXJbYarfd0bI8uGV3XI7Cz0rzD1jeaL1bFlh3N5jQHpbsRxcg5XmRfAeBT2esoSNQh_pIIvkwZduJiGPx_Vrb4nlNfLQie7u7haB34WQ",
    "e": "AQAB",
    "use": "enc",
    "qi": "G9rq715HxCz3mBJPhjJvRdijwUD8VpV1-F_D9omKp7mCYavDnXeT0JzSbSFLdVvqOIazaPWmhBhXDuFiNtHU_Qs8PDO70e1JYObpYCmThtPb9sP3CsBFCettpuMjoGO2SdAxID6mrGFU3XujpjTFEf2av-QH9VHAoLRCjsuRoZw",
    "dp": "W_6N77yxxESyGtUx2sZntDdPWkoapPg74DCkz2XXNtslaw1qEJY9TksWQ3GqFrk8E0MxsDE6MJHjtqAxxp3ioLAy7X-liQKhP-lmvMzXpnAF8v3M1w8Kzq57wEmtv_omB6Vqu5UZlH6kIIBqsnLlTWJ7nYnlOWT6EeQBstE7FNs",
    "alg": "RSA1_5",
    "dq": "lm8Ghx1Ng44Q2bnO2ukzoOlVg0vEZzSxuOmq4qPzZp5aYRWlF7JpyTbgLRcQ1vNpbCd2xNOZbYKQrm6psgNagaZK2pFWI5yaJjPMIirQR43wExh4DBJRYQkYvDxkbAQv64hhk8tLNEuG30zzSN0IxYwqBOOcpGKiZHlfYF1x-hU",
    "n": "uoOMm38LZqcMJbPvrMcjmyIzTHVR1fBaXP6siNjKSXo_SgbX3xiaGTShtPmrb6JIPIhrG43TVvYW5flEWzEM6aJCS7IUULAHPo4AxuLOfze9AH36fcVaV8CGTHT2H7m4ROXOaDSQNUB3iuqI8Qy80x8WIF1zSNzr7--rUp0itKkJpaap04b40Gjl5_Dnr45m7TBw7bGgWLQBdUZ5FYXg3QNXCrIgvTxSFHU_5bTN5nn9cF6Lz2t3Df6jRlOGZ30NMPCyswqXwzLJqKro6Gk6CamBl3_lToB0kcXfjeB2xnkaORqNzcPLntBXXDsOGalEEUE7sU2J-ansqwEmQUHbVw"
}`

	signer, err := Arweave([]byte(walletJWK))
	assert.NoError(t, err)

	data := []byte("test")
	sig, err := signer.Sign(data)
	assert.NoError(t, err)

	assert.True(t, signer.Verify(data, sig))
}

func TestArweaveSigner_PublicKey(t *testing.T) {
	var walletJWK = `{
    "p": "5sbDW0jFyxb2ewKkhpAhB2N8Rv-gglD3zWcgCJIcF0Q2vUlWUd07YKvhc6hLQq0FDScoD_E1CkywqYXFSsXu6oTP_dSKpK8wmUhzYDt3CqD2fr0FmoogQXhirxVuTfDKvX0MnwRpQeGR_M-TODGKoTs5tO89xqhwt5N2NsQi-8M",
    "kty": "RSA",
    "q": "zuZMWm1f5YBHYyPtptnMv3oyiUAFbmy7RNKNx3ck0fDuL_9CcpctKeniSlxRoNWxw808XwSmA1HyOERWacnAmbKhJOztKgK_14CqUMNPec86Nyfq6--h-sfUk6O0rmQm34aGflTIIVhStQQ7MSocvBTzNAeYE403s5pwfx2DLN0",
    "d": "pmvs0W6J0L2rNBUa_dLurWTV2xovvDJrnuDLOIwf63I9-KKkifdcwLLauy2vz7tAs2qSD5L8CXMpuWgymx_oQqu6depfACQhQqrYfJ4x3QfyB_dgxEqNMxGhg6U1W9eOWUOoagjEyA9O21AsgtF3xvXM-3IHBPPDPmNfPibhqcpIPd8L0Q9Vr1T65YrNsIUDbPKgWcwpk6hAzq-H31_Gjy4AteWFF6KINj7TAfMh_IRRNApXJbYarfd0bI8uGV3XI7Cz0rzD1jeaL1bFlh3N5jQHpbsRxcg5XmRfAeBT2esoSNQh_pIIvkwZduJiGPx_Vrb4nlNfLQie7u7haB34WQ",
    "e": "AQAB",
    "use": "enc",
    "qi": "G9rq715HxCz3mBJPhjJvRdijwUD8VpV1-F_D9omKp7mCYavDnXeT0JzSbSFLdVvqOIazaPWmhBhXDuFiNtHU_Qs8PDO70e1JYObpYCmThtPb9sP3CsBFCettpuMjoGO2SdAxID6mrGFU3XujpjTFEf2av-QH9VHAoLRCjsuRoZw",
    "dp": "W_6N77yxxESyGtUx2sZntDdPWkoapPg74DCkz2XXNtslaw1qEJY9TksWQ3GqFrk8E0MxsDE6MJHjtqAxxp3ioLAy7X-liQKhP-lmvMzXpnAF8v3M1w8Kzq57wEmtv_omB6Vqu5UZlH6kIIBqsnLlTWJ7nYnlOWT6EeQBstE7FNs",
    "alg": "RSA1_5",
    "dq": "lm8Ghx1Ng44Q2bnO2ukzoOlVg0vEZzSxuOmq4qPzZp5aYRWlF7JpyTbgLRcQ1vNpbCd2xNOZbYKQrm6psgNagaZK2pFWI5yaJjPMIirQR43wExh4DBJRYQkYvDxkbAQv64hhk8tLNEuG30zzSN0IxYwqBOOcpGKiZHlfYF1x-hU",
    "n": "uoOMm38LZqcMJbPvrMcjmyIzTHVR1fBaXP6siNjKSXo_SgbX3xiaGTShtPmrb6JIPIhrG43TVvYW5flEWzEM6aJCS7IUULAHPo4AxuLOfze9AH36fcVaV8CGTHT2H7m4ROXOaDSQNUB3iuqI8Qy80x8WIF1zSNzr7--rUp0itKkJpaap04b40Gjl5_Dnr45m7TBw7bGgWLQBdUZ5FYXg3QNXCrIgvTxSFHU_5bTN5nn9cF6Lz2t3Df6jRlOGZ30NMPCyswqXwzLJqKro6Gk6CamBl3_lToB0kcXfjeB2xnkaORqNzcPLntBXXDsOGalEEUE7sU2J-ansqwEmQUHbVw"
}`

	signer, err := Arweave([]byte(walletJWK))
	assert.NoError(t, err)

	expected := sha256.Sum256(signer.publicKey.N.Bytes())
	assert.Equal(t, expected[:], signer.PublicKey())
}
