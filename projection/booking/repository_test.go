package booking

import (
	"reflect"
	"testing"

	rp "github.com/10xLabs/chandler/repository"
	"github.com/10xLabs/chandler/store"
)

func TestNewRepository(t *testing.T) {
	d := ""
	type args struct {
		s   store.Store
		dir *string
	}
	tests := []struct {
		name string
		args args
		want Repository
	}{
		{
			name: "NewRepository",
			args: args{
				s: nil, dir: &d,
			},
			want: &repository{&rp.Base{Store: nil, CollectionName: &d}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(tt.args.s, tt.args.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
