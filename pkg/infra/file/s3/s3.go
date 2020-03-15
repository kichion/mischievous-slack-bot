package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3 S3の情報を保持する構造体です
type S3 struct {
	*s3.S3
	Az     string
	Bucket string
}

// Selecable はS3を閲覧する振る舞いを表現するインターフェースです
type Selecable interface {
	Region() string
	S3() string
}

// NewClient はs3を取り回すClientを生成して返します
func NewClient(v Selecable) (*S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(v.Region()),
	})
	if err != nil {
		return nil, err
	}
	return &S3{
		s3.New(sess),
		v.Region(),
		v.S3(),
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
