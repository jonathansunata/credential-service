package util

import (
	"crypto/rsa"
	"fmt"
	error_handler "github.com/SawitProRecruitment/UserService/error"
	"github.com/golang-jwt/jwt/v4"
)

func GeneratePrivateKey(filePath string) (*rsa.PrivateKey, error) {
	privateKeyFile := "-----BEGIN PRIVATE KEY-----\nMIIEuwIBADANBgkqhkiG9w0BAQEFAASCBKUwggShAgEAAoIBAQC0DZ7ChAVSNVoA\nLejYJGpj74q5pcY4sbaqLPNZv4eQR8dIckbvC0omSHe+3UBJeBGLxOckct6TS24q\n+2xEfj7igCwsaaPu8Iu5lMSKwpG/lgzu/t32bxXEIEPSGOxiD7CGN2/E5DtmIIVj\nMRLgT2ELuQGXtpduIkW92o+7yCWcbXcyEg0PQe/ImTgnyUo7vGMbX7WvDPb5It0M\ns7esgWA5t1dzH/ul5XVXIH8hFB5uv2fND9I3QrwseQDwcbSYLRjOGEIOiH4xgcfb\nFeU2z914i6ImjU3azeS5Y5QtJki5ZoID5NohPi2V+kH41xonM82X5/0jhCeVVOMl\n2FoERYNtAgMBAAECgf9bjJFXVh6z2MIU/+2PTkofhizYjZxm4fV+5wbVG2L9JCyv\n1BXnywBozBrBcWi+nSOkGE9uuPly62eYyJNu6sIOvh0NIqm3/uxh8Bf/IJn6+tCY\ntKkneals5PrsWL/Yccx6cgrikdRMmTyOsuUQ00uvf4zPwdzn+C/TLiOR111IVi9H\nVRmEytDlHTDjj8vvxs+ZYS8peqgNSBycox8MGjHLBJFP8/ntzouh/6MM6W4FApmS\n5+YvRECSb8VcaQJMIqcbDqt4RTJMItblqfYurV7UYT9GbS3pBQ8PoTnF2SkDxtwH\ns006Z8LhDEl+7rodQmm3XA0L877QUfGwW8F2pNECgYEA8VgDHt2qp/7nwR7+Ezne\naXeYvoqmIDQv08/Wm0VimKaVurw7eMrwbLVShW+kGoeMpP+htWtaW8u7Zc2chbOR\nrvWSsx8rb/cJlEWDhIK2z9gr+YMhFwlx1HKm8lztt67s+V5NRN28Oqxtwla9YRAj\n8TFHKLSGS0gatXOkD+UbrVUCgYEAvvzFFyG99NFesbO6dfH9SI1w2zLUrRRlHHVo\nndfEKcnIErklgCMrKysPVGSj2B1mpM8Ox7jl4NvxtHIg1FgIAh6VE/ZdYlYBRjPi\nopAYnyZF8MYneApUldqGT7sY5zCfJgke9hnMtGbPXhzqaVilRvqULR6tomqZYnOo\nby8YPbkCgYA1ZLZUBtBxmEhnhlbJpBzbknT9eqkkKMeIAcxFz8TvZrNre6dgou0r\n77WRBdD1eZWZD2EUROrZsioEbMe7IK4TWgsZi8TNYYcCAZsGHvEY7IdWDTet5A4F\n5VOf/QUuhQmyZbWMjc3N4UXrH8uIBM0e2DsY+09Wql4WVL4wMgy8fQKBgQCrLH2w\n98r6y1QlzMIHx/WMu0g1Dd/TqH3e/dPf9GyaT4GEVnCn4d1k+Vjp+LFolyFSAUpr\n8uoFmNuPMOL/rk6vJ53RoHOeGRtXQlWUAbYvnev9mnvxeMDK9mp+t1/ghZF+U5pu\nVD1GSwb8gMoP1SV88kUwE1joQsZqmOKTlBAT8QKBgAjpf8bRyG0H/LHyPD09DpH1\nXWw5exGr0oMqP7wEjxDGmXsmgvhbjNBlVuLPi54Sds9TD4Oo1x3h3lrE3xQ1lO+r\nslT5M4gdVL4w6Pi/pXjW5XKxtJqT+tJL+4Sn22lQrSyma++tH9eTj8ZO4/9PRv8Y\nEwGhILFGuqd8JHs9TQwN\n-----END PRIVATE KEY-----\n"
	fmt.Printf("File path : %v \n", filePath)
	fmt.Printf("Private key file : %v \n", privateKeyFile)

	bytePrivateKey := []byte(privateKeyFile)

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(bytePrivateKey)
	if err != nil {
		fmt.Errorf("Error when parse private key : %v \n", err)
		return nil, error_handler.NewCustomError(500, err.Error())
	}

	return privateKey, nil

}

func GeneratePublicKey(filePath string) (*rsa.PublicKey, error) {
	publicKeyFile := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtA2ewoQFUjVaAC3o2CRq\nY++KuaXGOLG2qizzWb+HkEfHSHJG7wtKJkh3vt1ASXgRi8TnJHLek0tuKvtsRH4+\n4oAsLGmj7vCLuZTEisKRv5YM7v7d9m8VxCBD0hjsYg+whjdvxOQ7ZiCFYzES4E9h\nC7kBl7aXbiJFvdqPu8glnG13MhIND0HvyJk4J8lKO7xjG1+1rwz2+SLdDLO3rIFg\nObdXcx/7peV1VyB/IRQebr9nzQ/SN0K8LHkA8HG0mC0YzhhCDoh+MYHH2xXlNs/d\neIuiJo1N2s3kuWOULSZIuWaCA+TaIT4tlfpB+NcaJzPNl+f9I4QnlVTjJdhaBEWD\nbQIDAQAB\n-----END PUBLIC KEY-----\n"

	bytePublicKey := []byte(publicKeyFile)
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(bytePublicKey)
	if err != nil {
		return nil, error_handler.NewCustomError(500, err.Error())
	}
	return publicKey, nil
}
