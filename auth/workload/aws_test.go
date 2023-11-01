package workload

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	wipResourceName = "iam/project/7b6b7db5-5e04-498c-a9a5-9a883b9462a4/service-principal/test/workload-identity-provider/aws"
	accessKeyID     = "ASIASDOME5YRY6JUQQ7O"
	secretAccessKey = "9mpg9prjH54qmpVB99l1YS8mMraFs8WvrZJC+C2o"
	securityToken   = "IQoJb3JpZ2luX2VjEOX//////////wEaCXVzLWVhc3QtMSJIMEYCIQDVVsx2ozIeu2VgW9AYbwjrR3hhOYuncJCPWHYjeKh6XwIhAKPK3WK77+Q+z//E2d1R332IVAJN3svmSO+oBi1M2t3tKrsFCH4QARoMMTQ0ODQ2NDE3NDQzIgzrrhKJVjIO3x6lvPIqmAWSplqdzMuuhOrTIhGASwxI2tMCZlTIDKTV9upDfUfecMfV2swtd9EATvhWMztoKzNfG+r175/x7U8C9Bsb5WMr9MAx0v821HSRcbWv3WPP5F3mq/A9sI4I5j4jJx8NFwxyF96k7qhfffnVjG37oyuNNAcL18IsuexHEFiVAnqaj4plaBP3ea6en2ANksHcNn3UgIsat8bphpHJxwZ5O3Jh3sIwWu/KNXSrPyWzGC4/N/SpkThKNckL6MDSJZj2fW9WRO56TSynz2mh2CQGF5rIOjpflk2YKll65xAHTj863ELglo89950rB4K00lRSB0k8xAeS4uoJcPk9HJ2PPcY0LH6gkONR+QH2y28TRDfDPVZcJWzGyw/PZl0qWcgXEo43CzkmLD36Fa4F8P05ipeSPW0z3G+xUNyUqupm8VN4QD8XsUpmTICobBQhFLs3MCRFwSQexjnzkjDDTJz/HNTfJGKsOD09o7gRdlLVKxrQw/JM/A2dsaDuMW6KqjmrOsE804eUUnucBnNRObfeMBUz+QneSynMSIysIAnZ3STUBRfHjuj/YgIg3uaN4h84PvFfrB9WOKwuG0QCkJFoIBtborg0y1vWnqf/Kp8EAy3lwr+IE6/t/5dHPtCp6uUxmT8O/yW4PHA2/RufQ8isazIt/a/IdVHMht8VaYJRv5d9OyiNL5Ohh9WkesJ3SPgv2tEpJU4z9b1IiN18i0Tk30H1+nuQKyQsTjX2nB6RnKIpNmNseQKIuMFvOt3MqvquseWtMKjiApUauFBRe4Xh3tzA0w0jSSOPfY6tOO7p6D2j1vsF8X2Bt+l6/NUNgyGE6mgWm0vAWkK0nk2n7t/1DSLTB0NemS5iQcYhl8x9sadIUGTtah/9pLEKMLDJkKYGOrABoU0ILA/R+ATv9zfk3sL8+FdhzckPg6XHRHlZPHT7CC+sdppijioSmnecS9mF/7kf/Qrv73UXGY6drdmV7TEXW3GqoINiN6VnLhVwXlBmKtKMm0njDnfPX1MHrY++ZObt+nqAlhZQy6LP6CL1U38Pk5yXA1+j14J4ZY1E+25heKQY95Zn3yHGnXxYy8A1ZKo+9jbihWhDuuZIrizN5QNid+QYtxEFJ7eod2U84bll21o="
)

func testNow() time.Time {
	return time.Date(2023, 7, 28, 15, 25, 4, 3, time.UTC)
}

func TestAWSCredentialSource_getCallerID(t *testing.T) {
	type env struct {
		region          bool
		accessKeyID     bool
		secretAccessKey bool
		securityToken   bool
	}

	tests := []struct {
		name string

		// whether we are using v2
		v2 bool

		// env controls whether the values are passed via env vars. If not they
		// are returned from the test imds server
		env env

		// values to be returned
		region             string
		rolename           string
		accessKeyID        string
		secretAccessKey    string
		securityToken      string
		imdsv2SessionToken string
		wip                string

		// expected response
		want *callerIdentityRequest
	}{
		{
			name: "with token env",
			env: env{
				region:          true,
				accessKeyID:     true,
				secretAccessKey: true,
				securityToken:   true,
			},
			v2:              true,
			region:          "us-east-1",
			accessKeyID:     accessKeyID,
			secretAccessKey: secretAccessKey,
			securityToken:   securityToken,
			wip:             wipResourceName,
			want: &callerIdentityRequest{
				Headers: map[string]string{
					"Authorization":                    "AWS4-HMAC-SHA256 Credential=ASIASDOME5YRY6JUQQ7O/20230728/us-east-1/sts/aws4_request, SignedHeaders=host;x-amz-date;x-amz-security-token;x-hcp-workload-identity-provider, Signature=4678f0beee2b89029de438cc2b75625b7074a5031304bb840c19f494f33bff6f",
					"Host":                             "sts.us-east-1.amazonaws.com",
					"X-Amz-Date":                       "20230728T152504Z",
					"X-Amz-Security-Token":             "IQoJb3JpZ2luX2VjEOX//////////wEaCXVzLWVhc3QtMSJIMEYCIQDVVsx2ozIeu2VgW9AYbwjrR3hhOYuncJCPWHYjeKh6XwIhAKPK3WK77+Q+z//E2d1R332IVAJN3svmSO+oBi1M2t3tKrsFCH4QARoMMTQ0ODQ2NDE3NDQzIgzrrhKJVjIO3x6lvPIqmAWSplqdzMuuhOrTIhGASwxI2tMCZlTIDKTV9upDfUfecMfV2swtd9EATvhWMztoKzNfG+r175/x7U8C9Bsb5WMr9MAx0v821HSRcbWv3WPP5F3mq/A9sI4I5j4jJx8NFwxyF96k7qhfffnVjG37oyuNNAcL18IsuexHEFiVAnqaj4plaBP3ea6en2ANksHcNn3UgIsat8bphpHJxwZ5O3Jh3sIwWu/KNXSrPyWzGC4/N/SpkThKNckL6MDSJZj2fW9WRO56TSynz2mh2CQGF5rIOjpflk2YKll65xAHTj863ELglo89950rB4K00lRSB0k8xAeS4uoJcPk9HJ2PPcY0LH6gkONR+QH2y28TRDfDPVZcJWzGyw/PZl0qWcgXEo43CzkmLD36Fa4F8P05ipeSPW0z3G+xUNyUqupm8VN4QD8XsUpmTICobBQhFLs3MCRFwSQexjnzkjDDTJz/HNTfJGKsOD09o7gRdlLVKxrQw/JM/A2dsaDuMW6KqjmrOsE804eUUnucBnNRObfeMBUz+QneSynMSIysIAnZ3STUBRfHjuj/YgIg3uaN4h84PvFfrB9WOKwuG0QCkJFoIBtborg0y1vWnqf/Kp8EAy3lwr+IE6/t/5dHPtCp6uUxmT8O/yW4PHA2/RufQ8isazIt/a/IdVHMht8VaYJRv5d9OyiNL5Ohh9WkesJ3SPgv2tEpJU4z9b1IiN18i0Tk30H1+nuQKyQsTjX2nB6RnKIpNmNseQKIuMFvOt3MqvquseWtMKjiApUauFBRe4Xh3tzA0w0jSSOPfY6tOO7p6D2j1vsF8X2Bt+l6/NUNgyGE6mgWm0vAWkK0nk2n7t/1DSLTB0NemS5iQcYhl8x9sadIUGTtah/9pLEKMLDJkKYGOrABoU0ILA/R+ATv9zfk3sL8+FdhzckPg6XHRHlZPHT7CC+sdppijioSmnecS9mF/7kf/Qrv73UXGY6drdmV7TEXW3GqoINiN6VnLhVwXlBmKtKMm0njDnfPX1MHrY++ZObt+nqAlhZQy6LP6CL1U38Pk5yXA1+j14J4ZY1E+25heKQY95Zn3yHGnXxYy8A1ZKo+9jbihWhDuuZIrizN5QNid+QYtxEFJ7eod2U84bll21o=",
					"X-Hcp-Workload-Identity-Provider": "iam/project/7b6b7db5-5e04-498c-a9a5-9a883b9462a4/service-principal/test/workload-identity-provider/aws",
				},
				Method: "POST",
				URL:    "https://sts.us-east-1.amazonaws.com?Action=GetCallerIdentity&Version=2011-06-15",
			},
		},
		{
			name: "without token env",
			env: env{
				region:          true,
				accessKeyID:     true,
				secretAccessKey: true,
				securityToken:   false,
			},
			region:          "us-east-1",
			accessKeyID:     accessKeyID,
			secretAccessKey: secretAccessKey,
			wip:             wipResourceName,
			want: &callerIdentityRequest{
				Headers: map[string]string{
					"Authorization":                    "AWS4-HMAC-SHA256 Credential=ASIASDOME5YRY6JUQQ7O/20230728/us-east-1/sts/aws4_request, SignedHeaders=host;x-amz-date;x-hcp-workload-identity-provider, Signature=30deb4075f257d5a8e77a8b8b4d2fa962386e4c88fad954c56367df201f2943d",
					"Host":                             "sts.us-east-1.amazonaws.com",
					"X-Amz-Date":                       "20230728T152504Z",
					"X-Hcp-Workload-Identity-Provider": "iam/project/7b6b7db5-5e04-498c-a9a5-9a883b9462a4/service-principal/test/workload-identity-provider/aws",
				},
				Method: "POST",
				URL:    "https://sts.us-east-1.amazonaws.com?Action=GetCallerIdentity&Version=2011-06-15",
			},
		},
		{
			name:            "without token all metadata",
			region:          "us-east-1",
			accessKeyID:     accessKeyID,
			secretAccessKey: secretAccessKey,
			wip:             wipResourceName,
			want: &callerIdentityRequest{
				Headers: map[string]string{
					"Authorization":                    "AWS4-HMAC-SHA256 Credential=ASIASDOME5YRY6JUQQ7O/20230728/us-east-1/sts/aws4_request, SignedHeaders=host;x-amz-date;x-hcp-workload-identity-provider, Signature=30deb4075f257d5a8e77a8b8b4d2fa962386e4c88fad954c56367df201f2943d",
					"Host":                             "sts.us-east-1.amazonaws.com",
					"X-Amz-Date":                       "20230728T152504Z",
					"X-Hcp-Workload-Identity-Provider": "iam/project/7b6b7db5-5e04-498c-a9a5-9a883b9462a4/service-principal/test/workload-identity-provider/aws",
				},
				Method: "POST",
				URL:    "https://sts.us-east-1.amazonaws.com?Action=GetCallerIdentity&Version=2011-06-15",
			},
		},
		{
			name:            "with token all metadata v1",
			region:          "us-east-1",
			accessKeyID:     accessKeyID,
			secretAccessKey: secretAccessKey,
			securityToken:   securityToken,
			wip:             wipResourceName,
			want: &callerIdentityRequest{
				Headers: map[string]string{
					"Authorization":                    "AWS4-HMAC-SHA256 Credential=ASIASDOME5YRY6JUQQ7O/20230728/us-east-1/sts/aws4_request, SignedHeaders=host;x-amz-date;x-hcp-workload-identity-provider, Signature=30deb4075f257d5a8e77a8b8b4d2fa962386e4c88fad954c56367df201f2943d",
					"Host":                             "sts.us-east-1.amazonaws.com",
					"X-Amz-Date":                       "20230728T152504Z",
					"X-Hcp-Workload-Identity-Provider": "iam/project/7b6b7db5-5e04-498c-a9a5-9a883b9462a4/service-principal/test/workload-identity-provider/aws",
				},
				Method: "POST",
				URL:    "https://sts.us-east-1.amazonaws.com?Action=GetCallerIdentity&Version=2011-06-15",
			},
		},
		{
			name:               "with token all metadata v2",
			v2:                 true,
			imdsv2SessionToken: "test-session-token",
			region:             "us-east-1",
			accessKeyID:        accessKeyID,
			secretAccessKey:    secretAccessKey,
			securityToken:      securityToken,
			wip:                wipResourceName,
			want: &callerIdentityRequest{
				Headers: map[string]string{
					"Authorization":                    "AWS4-HMAC-SHA256 Credential=ASIASDOME5YRY6JUQQ7O/20230728/us-east-1/sts/aws4_request, SignedHeaders=host;x-amz-date;x-amz-security-token;x-hcp-workload-identity-provider, Signature=4678f0beee2b89029de438cc2b75625b7074a5031304bb840c19f494f33bff6f",
					"Host":                             "sts.us-east-1.amazonaws.com",
					"X-Amz-Date":                       "20230728T152504Z",
					"X-Amz-Security-Token":             "IQoJb3JpZ2luX2VjEOX//////////wEaCXVzLWVhc3QtMSJIMEYCIQDVVsx2ozIeu2VgW9AYbwjrR3hhOYuncJCPWHYjeKh6XwIhAKPK3WK77+Q+z//E2d1R332IVAJN3svmSO+oBi1M2t3tKrsFCH4QARoMMTQ0ODQ2NDE3NDQzIgzrrhKJVjIO3x6lvPIqmAWSplqdzMuuhOrTIhGASwxI2tMCZlTIDKTV9upDfUfecMfV2swtd9EATvhWMztoKzNfG+r175/x7U8C9Bsb5WMr9MAx0v821HSRcbWv3WPP5F3mq/A9sI4I5j4jJx8NFwxyF96k7qhfffnVjG37oyuNNAcL18IsuexHEFiVAnqaj4plaBP3ea6en2ANksHcNn3UgIsat8bphpHJxwZ5O3Jh3sIwWu/KNXSrPyWzGC4/N/SpkThKNckL6MDSJZj2fW9WRO56TSynz2mh2CQGF5rIOjpflk2YKll65xAHTj863ELglo89950rB4K00lRSB0k8xAeS4uoJcPk9HJ2PPcY0LH6gkONR+QH2y28TRDfDPVZcJWzGyw/PZl0qWcgXEo43CzkmLD36Fa4F8P05ipeSPW0z3G+xUNyUqupm8VN4QD8XsUpmTICobBQhFLs3MCRFwSQexjnzkjDDTJz/HNTfJGKsOD09o7gRdlLVKxrQw/JM/A2dsaDuMW6KqjmrOsE804eUUnucBnNRObfeMBUz+QneSynMSIysIAnZ3STUBRfHjuj/YgIg3uaN4h84PvFfrB9WOKwuG0QCkJFoIBtborg0y1vWnqf/Kp8EAy3lwr+IE6/t/5dHPtCp6uUxmT8O/yW4PHA2/RufQ8isazIt/a/IdVHMht8VaYJRv5d9OyiNL5Ohh9WkesJ3SPgv2tEpJU4z9b1IiN18i0Tk30H1+nuQKyQsTjX2nB6RnKIpNmNseQKIuMFvOt3MqvquseWtMKjiApUauFBRe4Xh3tzA0w0jSSOPfY6tOO7p6D2j1vsF8X2Bt+l6/NUNgyGE6mgWm0vAWkK0nk2n7t/1DSLTB0NemS5iQcYhl8x9sadIUGTtah/9pLEKMLDJkKYGOrABoU0ILA/R+ATv9zfk3sL8+FdhzckPg6XHRHlZPHT7CC+sdppijioSmnecS9mF/7kf/Qrv73UXGY6drdmV7TEXW3GqoINiN6VnLhVwXlBmKtKMm0njDnfPX1MHrY++ZObt+nqAlhZQy6LP6CL1U38Pk5yXA1+j14J4ZY1E+25heKQY95Zn3yHGnXxYy8A1ZKo+9jbihWhDuuZIrizN5QNid+QYtxEFJ7eod2U84bll21o=",
					"X-Hcp-Workload-Identity-Provider": "iam/project/7b6b7db5-5e04-498c-a9a5-9a883b9462a4/service-principal/test/workload-identity-provider/aws",
				},
				Method: "POST",
				URL:    "https://sts.us-east-1.amazonaws.com?Action=GetCallerIdentity&Version=2011-06-15",
			},
		},
		{
			name: "region env var v1",
			env: env{
				region: true,
			},
			rolename:        "test-role",
			region:          "us-east-1",
			accessKeyID:     accessKeyID,
			secretAccessKey: secretAccessKey,
			wip:             wipResourceName,
			want: &callerIdentityRequest{
				Headers: map[string]string{
					"Authorization":                    "AWS4-HMAC-SHA256 Credential=ASIASDOME5YRY6JUQQ7O/20230728/us-east-1/sts/aws4_request, SignedHeaders=host;x-amz-date;x-hcp-workload-identity-provider, Signature=30deb4075f257d5a8e77a8b8b4d2fa962386e4c88fad954c56367df201f2943d",
					"Host":                             "sts.us-east-1.amazonaws.com",
					"X-Amz-Date":                       "20230728T152504Z",
					"X-Hcp-Workload-Identity-Provider": "iam/project/7b6b7db5-5e04-498c-a9a5-9a883b9462a4/service-principal/test/workload-identity-provider/aws",
				},
				Method: "POST",
				URL:    "https://sts.us-east-1.amazonaws.com?Action=GetCallerIdentity&Version=2011-06-15",
			},
		},
		{
			name: "region env var v2",
			env: env{
				region: true,
			},
			v2:                 true,
			imdsv2SessionToken: "test-session-token",
			rolename:           "test-role",
			region:             "us-east-1",
			accessKeyID:        accessKeyID,
			secretAccessKey:    secretAccessKey,
			securityToken:      securityToken,
			wip:                wipResourceName,
			want: &callerIdentityRequest{
				Headers: map[string]string{
					"Authorization":                    "AWS4-HMAC-SHA256 Credential=ASIASDOME5YRY6JUQQ7O/20230728/us-east-1/sts/aws4_request, SignedHeaders=host;x-amz-date;x-amz-security-token;x-hcp-workload-identity-provider, Signature=4678f0beee2b89029de438cc2b75625b7074a5031304bb840c19f494f33bff6f",
					"Host":                             "sts.us-east-1.amazonaws.com",
					"X-Amz-Date":                       "20230728T152504Z",
					"X-Amz-Security-Token":             "IQoJb3JpZ2luX2VjEOX//////////wEaCXVzLWVhc3QtMSJIMEYCIQDVVsx2ozIeu2VgW9AYbwjrR3hhOYuncJCPWHYjeKh6XwIhAKPK3WK77+Q+z//E2d1R332IVAJN3svmSO+oBi1M2t3tKrsFCH4QARoMMTQ0ODQ2NDE3NDQzIgzrrhKJVjIO3x6lvPIqmAWSplqdzMuuhOrTIhGASwxI2tMCZlTIDKTV9upDfUfecMfV2swtd9EATvhWMztoKzNfG+r175/x7U8C9Bsb5WMr9MAx0v821HSRcbWv3WPP5F3mq/A9sI4I5j4jJx8NFwxyF96k7qhfffnVjG37oyuNNAcL18IsuexHEFiVAnqaj4plaBP3ea6en2ANksHcNn3UgIsat8bphpHJxwZ5O3Jh3sIwWu/KNXSrPyWzGC4/N/SpkThKNckL6MDSJZj2fW9WRO56TSynz2mh2CQGF5rIOjpflk2YKll65xAHTj863ELglo89950rB4K00lRSB0k8xAeS4uoJcPk9HJ2PPcY0LH6gkONR+QH2y28TRDfDPVZcJWzGyw/PZl0qWcgXEo43CzkmLD36Fa4F8P05ipeSPW0z3G+xUNyUqupm8VN4QD8XsUpmTICobBQhFLs3MCRFwSQexjnzkjDDTJz/HNTfJGKsOD09o7gRdlLVKxrQw/JM/A2dsaDuMW6KqjmrOsE804eUUnucBnNRObfeMBUz+QneSynMSIysIAnZ3STUBRfHjuj/YgIg3uaN4h84PvFfrB9WOKwuG0QCkJFoIBtborg0y1vWnqf/Kp8EAy3lwr+IE6/t/5dHPtCp6uUxmT8O/yW4PHA2/RufQ8isazIt/a/IdVHMht8VaYJRv5d9OyiNL5Ohh9WkesJ3SPgv2tEpJU4z9b1IiN18i0Tk30H1+nuQKyQsTjX2nB6RnKIpNmNseQKIuMFvOt3MqvquseWtMKjiApUauFBRe4Xh3tzA0w0jSSOPfY6tOO7p6D2j1vsF8X2Bt+l6/NUNgyGE6mgWm0vAWkK0nk2n7t/1DSLTB0NemS5iQcYhl8x9sadIUGTtah/9pLEKMLDJkKYGOrABoU0ILA/R+ATv9zfk3sL8+FdhzckPg6XHRHlZPHT7CC+sdppijioSmnecS9mF/7kf/Qrv73UXGY6drdmV7TEXW3GqoINiN6VnLhVwXlBmKtKMm0njDnfPX1MHrY++ZObt+nqAlhZQy6LP6CL1U38Pk5yXA1+j14J4ZY1E+25heKQY95Zn3yHGnXxYy8A1ZKo+9jbihWhDuuZIrizN5QNid+QYtxEFJ7eod2U84bll21o=",
					"X-Hcp-Workload-Identity-Provider": "iam/project/7b6b7db5-5e04-498c-a9a5-9a883b9462a4/service-principal/test/workload-identity-provider/aws",
				},
				Method: "POST",
				URL:    "https://sts.us-east-1.amazonaws.com?Action=GetCallerIdentity&Version=2011-06-15",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			// Set the env vars
			if tt.env.region {
				t.Setenv(awsEnvRegion, tt.region)
			}
			if tt.env.accessKeyID {
				t.Setenv(awsEnvAccessKeyID, tt.accessKeyID)
			}
			if tt.env.secretAccessKey {
				t.Setenv(awsEnvSecretAccessKey, tt.secretAccessKey)
			}
			if tt.env.securityToken {
				t.Setenv(awsEnvSessionToken, tt.securityToken)
			}

			mockAWS := createAwsTestServer(t, tt.rolename, tt.region, tt.accessKeyID, tt.secretAccessKey, tt.securityToken, tt.imdsv2SessionToken)
			client := mockAWS.server.Client()
			transport, ok := client.Transport.(*http.Transport)
			r.True(ok)
			transport.Proxy = func(req *http.Request) (*url.URL, error) {
				url, err := url.Parse(fmt.Sprintf("%s%s", mockAWS.server.URL, req.URL.Path))
				return url, err
			}

			// Create the AWS CredSource
			ac := &AWSCredentialSource{
				IMDSv2: tt.v2,
				now:    testNow,
				client: client,
			}
			out, err := ac.getCallerIdentityReq(tt.wip)
			r.NoError(err)
			r.Equal(tt.want.URL, out.URL)
			r.Equal(tt.want.Method, out.Method)
			r.EqualValues(tt.want.Headers, out.Headers)
		})
	}
}

type testAwsServer struct {
	t        *testing.T
	server   *httptest.Server
	rolename string
	region   string

	accessKey       string
	secretAccessKey string
	securityToken   string

	imdsv2SessionToken string
}

func createAwsTestServer(t *testing.T, rolename, region, accessKey, secretAccessKey, securityToken, imdsv2SessionToken string) *testAwsServer {
	aws := &testAwsServer{
		rolename:           rolename,
		region:             region,
		accessKey:          accessKey,
		secretAccessKey:    secretAccessKey,
		securityToken:      securityToken,
		imdsv2SessionToken: imdsv2SessionToken,
	}

	aws.server = httptest.NewTLSServer(aws)
	t.Cleanup(aws.server.Close)
	return aws
}

func (aws *testAwsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	validateSessionToken := func(r *http.Request) {
		if aws.imdsv2SessionToken != "" {
			headerValue := r.Header.Get(awsIMDSv2SessionTokenHeader)
			if headerValue != aws.imdsv2SessionToken {
				aws.t.Errorf("%q = \n%q\n want \n%q", awsIMDSv2SessionTokenHeader, headerValue, aws.imdsv2SessionToken)
			}
		}
	}

	validateSessionTTL := func(r *http.Request) {
		if aws.imdsv2SessionToken != "" {
			headerValue := r.Header.Get(awsIMDSv2SessionTTLHeader)
			if headerValue != awsIMDSv2SessionTTL {
				aws.t.Errorf("%q = \n%q\n want \n%q", awsIMDSv2SessionTTLHeader, headerValue, awsIMDSv2SessionTTL)
			}
		}
	}

	var err error
	switch p := r.URL.Path; p {
	case "/latest/meta-data/iam/security-credentials":
		validateSessionToken(r)
		_, err = w.Write([]byte(aws.rolename))
	case fmt.Sprintf("/latest/meta-data/iam/security-credentials/%s", aws.rolename):
		validateSessionToken(r)

		// Build the response
		creds := map[string]string{}
		if aws.accessKey != "" {
			creds["AccessKeyId"] = aws.accessKey
		}
		if aws.secretAccessKey != "" {
			creds["SecretAccessKey"] = aws.secretAccessKey
		}
		if aws.imdsv2SessionToken != "" && aws.securityToken != "" {
			creds["Token"] = aws.securityToken
		}

		jsonCredentials, _ := json.Marshal(creds)
		_, err = w.Write(jsonCredentials)
	case "/latest/meta-data/placement/region":
		validateSessionToken(r)
		_, err = w.Write([]byte(aws.region))
	case "/latest/api/token":
		validateSessionTTL(r)
		_, err = w.Write([]byte(aws.imdsv2SessionToken))
	}

	if err != nil {
		aws.t.Fatalf("unexpected error: %v", err)
	}
}
