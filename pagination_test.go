package pagination

import (
	"reflect"
	"testing"
)

// TestPagination_New 测试分页构造函数
func TestPagination_New(t *testing.T) {
	type args struct {
		page     int
		pageSize int
		total    int
	}
	tests := []struct {
		name string
		args args
		want *Pagination[string]
	}{
		{
			name: "正常情况",
			args: args{page: 1, pageSize: 10, total: 100},
			want: &Pagination[string]{
				Page:       1,
				PageSize:   10,
				TotalCount: 100,
				PageCount:  10,
				IsLastPage: false,
			},
		},
		{
			name: "页码大于总页数",
			args: args{page: 15, pageSize: 10, total: 100},
			want: &Pagination[string]{
				Page:       10,
				PageSize:   10,
				TotalCount: 100,
				PageCount:  10,
				IsLastPage: true,
			},
		},
		{
			name: "页面大小无效",
			args: args{page: 1, pageSize: 0, total: 100},
			want: &Pagination[string]{
				Page:       1,
				PageSize:   10,
				TotalCount: 100,
				PageCount:  10,
				IsLastPage: false,
			},
		},
		{
			name: "负总数",
			args: args{page: 1, pageSize: 10, total: -1},
			want: &Pagination[string]{
				Page:       1,
				PageSize:   10,
				TotalCount: -1,
				PageCount:  -1,
				IsLastPage: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New[string](tt.args.page, tt.args.pageSize, tt.args.total)
			if got.Page != tt.want.Page {
				t.Errorf("New().Page = %v, want %v", got.Page, tt.want.Page)
			}
			if got.PageSize != tt.want.PageSize {
				t.Errorf("New().PageSize = %v, want %v", got.PageSize, tt.want.PageSize)
			}
			if got.TotalCount != tt.want.TotalCount {
				t.Errorf("New().TotalCount = %v, want %v", got.TotalCount, tt.want.TotalCount)
			}
			if got.PageCount != tt.want.PageCount {
				t.Errorf("New().PageCount = %v, want %v", got.PageCount, tt.want.PageCount)
			}
			if got.IsLastPage != tt.want.IsLastPage {
				t.Errorf("New().IsLastPage = %v, want %v", got.IsLastPage, tt.want.IsLastPage)
			}
		})
	}
}

// TestPagination_SetItems 测试设置项目列表功能
func TestPagination_SetItems(t *testing.T) {
	p := New[string](1, 10, 100)
	items := []string{"item1", "item2", "item3"}
	result := p.SetItems(items)

	if !reflect.DeepEqual(p.Items, items) {
		t.Errorf("SetItems() 失败, got %v, want %v", p.Items, items)
	}

	// 检查是否返回相同实例
	if result != p {
		t.Error("SetItems() 应该返回相同实例")
	}
}

// TestPagination_AddItem 测试添加单个项目功能
func TestPagination_AddItem(t *testing.T) {
	p := New[string](1, 10, 100)
	p.SetItems([]string{"item1", "item2"})

	// 添加新项目
	result := p.AddItem("item3")

	// 检查项目是否已添加
	expected := []string{"item1", "item2", "item3"}
	if !reflect.DeepEqual(p.Items, expected) {
		t.Errorf("AddItem() 失败, got %v, want %v", p.Items, expected)
	}

	// 检查是否返回相同实例
	if result != p {
		t.Error("AddItem() 应该返回相同实例")
	}
}

// TestPagination_WithDifferentTypes 测试不同类型的分页
func TestPagination_WithDifferentTypes(t *testing.T) {
	// 测试整数类型
	intPagination := New[int](1, 5, 20)
	intItems := []int{1, 2, 3, 4, 5}
	intPagination.SetItems(intItems)

	if !reflect.DeepEqual(intPagination.Items, intItems) {
		t.Errorf("Pagination[int] 失败, got %v, want %v", intPagination.Items, intItems)
	}

	// 测试自定义结构体类型
	type Person struct {
		Name string
		Age  int
	}
	personPagination := New[Person](1, 3, 10)
	personItems := []Person{
		{Name: "张三", Age: 30},
		{Name: "李四", Age: 25},
	}
	personPagination.SetItems(personItems)

	if !reflect.DeepEqual(personPagination.Items, personItems) {
		t.Errorf("Pagination[Person] 失败, got %v, want %v", personPagination.Items, personItems)
	}
}
