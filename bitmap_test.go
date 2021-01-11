package bitmap

import (
	"testing"
)

func TestBitmap_Set(t *testing.T) {
	type fields struct {
		data uint64
	}
	type args struct {
		val      bool
		position uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    uint64
	}{
		{
			args: args{
				val:      true,
				position: 0,
			},
			fields: fields{
				data: 0,
			},
			wantErr: false,
			want:    1,
		},
		{
			args: args{
				val:      true,
				position: 2,
			},
			fields: fields{
				data: 0,
			},
			wantErr: false,
			want:    4,
		},
		{
			args: args{
				val:      true,
				position: 64,
			},
			fields: fields{
				data: 0,
			},
			wantErr: true,
			want:    4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bitmap{
				data: tt.fields.data,
			}
			if err := b.Set(tt.args.val, tt.args.position); (err != nil) != tt.wantErr {
				t.Errorf("Bitmap.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && b.data != tt.want {
				t.Errorf("Bitmap.Set() value = %v, want value %v", tt.fields.data, tt.want)
			}
		})
	}
}

func TestBitmap_Get(t *testing.T) {
	type fields struct {
		data uint64
	}
	type args struct {
		position uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			args: args{
				position: 2,
			},
			fields: fields{
				data: 4,
			},
			want:    true,
			wantErr: false,
		},
		{
			args: args{
				position: 1,
			},
			fields: fields{
				data: 4,
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bitmap{
				data: tt.fields.data,
			}
			got, err := b.Get(tt.args.position)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bitmap.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bitmap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
