package locache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCache(t *testing.T) {
	t.Parallel()

	t.Run("not nil", func(t *testing.T) {
		t.Parallel()
		require.NotNil(t, NewCache())
	})
}

func TestCache_Get(t *testing.T) {
	t.Parallel()

	type fields struct {
		m map[string]any
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
		want1  bool
	}{
		{
			name: "false on key not found",
		},
		{
			name: "true on key found but nil value",
			fields: fields{
				m: map[string]any{
					"key": nil,
				},
			},
			args:  args{"key"},
			want1: true,
		},
		{
			name: "true on key found and non-nil value",
			fields: fields{
				m: map[string]any{
					"key": "value",
				},
			},
			args:  args{"key"},
			want:  "value",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &Cache{
				m: tt.fields.m,
			}
			got, got1 := c.Get(tt.args.key)

			require.Equal(t, tt.want, got)
			require.Equal(t, tt.want1, got1)
		})
	}
}

func TestCache_Add(t *testing.T) {
	t.Parallel()

	type fields struct {
		m map[string]any
	}
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectMap map[string]any
	}{
		{
			name: "value added on key not found",
			fields: fields{
				m: map[string]any{},
			},
			args: args{
				key:   "key",
				value: "value",
			},
			expectMap: map[string]any{
				"key": "value",
			},
		},
		{
			name: "value not added on key found but empty value",
			fields: fields{
				m: map[string]any{
					"key": "",
				},
			},
			args: args{
				key:   "key",
				value: "value",
			},
			expectMap: map[string]any{
				"key": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &Cache{
				m: tt.fields.m,
			}
			c.Add(tt.args.key, tt.args.value)

			require.Equal(t, tt.expectMap, c.m)
		})
	}
}

func TestCache_Set(t *testing.T) {
	t.Parallel()

	type fields struct {
		m map[string]any
	}
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectMap map[string]any
	}{
		{
			name: "value added on key not found",
			fields: fields{
				m: map[string]any{},
			},
			args: args{
				key:   "key",
				value: "value",
			},
			expectMap: map[string]any{
				"key": "value",
			},
		},
		{
			name: "value replaced on key found with non-empty value",
			fields: fields{
				m: map[string]any{
					"key": "value1",
				},
			},
			args: args{
				key:   "key",
				value: "value2",
			},
			expectMap: map[string]any{
				"key": "value2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				m: tt.fields.m,
			}
			c.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestCache_Has(t *testing.T) {
	t.Parallel()

	type fields struct {
		m map[string]any
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "false on key not found",
			fields: fields{
				m: map[string]any{},
			},
		},
		{
			name: "true on key found with non-empty value",
			fields: fields{
				m: map[string]any{
					"key": "value",
				},
			},
			args: args{"key"},
			want: true,
		},
		{
			name: "true on key found with empty value",
			fields: fields{
				m: map[string]any{
					"key": "",
				},
			},
			args: args{"key"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &Cache{
				m: tt.fields.m,
			}
			got := c.Has(tt.args.key)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestCache_Delete(t *testing.T) {
	t.Parallel()

	type fields struct {
		m map[string]any
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "delete on key found",
			fields: fields{
				m: map[string]any{
					"key": "value",
				},
			},
			args: args{"key"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &Cache{
				m: tt.fields.m,
			}
			c.Delete(tt.args.key)

			require.NotContains(t, c.m, tt.args.key)
		})
	}
}
