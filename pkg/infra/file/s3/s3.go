package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
)

// S3 S3の情報を保持する構造体です
type S3 struct {
	*s3.S3
	Az     string
	Bucket string
}

// NewClient はs3を取り回すClientを生成して返します
func NewClient(v *environment.TalkMaster) (*S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(v.AWSRegion),
	})
	if err != nil {
		return nil, err
	}
	return &S3{
		s3.New(sess),
		v.AWSRegion,
		v.S3Storage,
	}, nil
}

func (s S3) buildParams(path string) *s3.SelectObjectContentInput {
	return &s3.SelectObjectContentInput{
		Bucket:         aws.String(s.Bucket),
		Key:            aws.String(path),
		ExpressionType: aws.String(s3.ExpressionTypeSql),
		Expression:     aws.String("SELECT * FROM S3Object"),
		InputSerialization: &s3.InputSerialization{
			CompressionType: aws.String("NONE"),
			JSON: &s3.JSONInput{
				Type: aws.String("Lines"),
			},
		},
		OutputSerialization: &s3.OutputSerialization{
			JSON: &s3.JSONOutput{},
		},
	}
}
