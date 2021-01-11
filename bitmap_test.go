package bitmap

import (
	"testing"
)

func Test_bitmap_Get(t *testing.T) {
	_bitmapTC1 := NewBitmap(100)
	_bitmapTC1.data[0] = 1

	_bitmapTC2 := NewBitmap(100)
	_bitmapTC2.data[0] = 2

	_bitmapTC3 := NewBitmap(100)
	_bitmapTC3.data[1] = 2

	type fields struct {
		bitmap *bitmap
	}
	type args struct {
		position uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			fields: fields{
				bitmap: _bitmapTC1,
			},
			args: args{
				position: 0,
			},
			want:    true,
			wantErr: false,
		},
		{
			fields: fields{
				bitmap: _bitmapTC2,
			},
			args: args{
				position: 1,
			},
			want:    true,
			wantErr: false,
		},
		{
			fields: fields{
				bitmap: _bitmapTC3,
			},
			args: args{
				position: 9,
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.bitmap.Get(tt.args.position)
			if (err != nil) != tt.wantErr {
				t.Errorf("bitmap.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("bitmap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitmap_Set(t *testing.T) {
	type fields struct {
		bitmap *bitmap
	}
	type args struct {
		val      bool
		position uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			fields: fields{
				bitmap: NewBitmap(100),
			},
			args: args{
				val:      true,
				position: 1,
			},
			wantErr: false,
		},
		{
			fields: fields{
				bitmap: NewBitmap(100),
			},
			args: args{
				val:      false,
				position: 1,
			},
			wantErr: false,
		},
		{
			fields: fields{
				bitmap: NewBitmap(100),
			},
			args: args{
				val:      true,
				position: 10,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.bitmap.Set(tt.args.position, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("bitmap.Set() error = %v, wantErr %v", err, tt.wantErr)
			}

			val, err := tt.fields.bitmap.Get(tt.args.position)
			if !tt.wantErr && err == nil && val != tt.args.val {
				t.Errorf("bitmap.Set() value = %v, want %v", val, tt.args.val)
			}
		})
	}
}
