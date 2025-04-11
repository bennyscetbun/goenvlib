package goenvlib

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetenvString(t *testing.T) {
	var addr *string
	var otherAddr *string
	// Test default value
	{
		t.Setenv("TEST_STRING", "")
		addr = GetenvString("TEST_STRING", "default")
		assert.Equal(t, "default", *addr)
	}

	// Test if the address is the same
	{
		str := GetenvString("TEST_STRING", "bla")
		assert.Equal(t, "default", *str)
		assert.Equal(t, addr, str)
	}
	// Test other env
	{
		t.Setenv("TEST_OTHER_STRING", "bla")
		otherAddr = GetenvString("TEST_OTHER_STRING", "default")
		assert.Equal(t, "bla", *otherAddr)
		assert.Assert(t, addr != otherAddr)
	}

	// Test if the address is the same
	{
		t.Setenv("TEST_STRING", "Pouic")
		t.Setenv("TEST_OTHER_STRING", "blou")
		assert.Equal(t, "default", *addr)
		assert.Equal(t, "bla", *otherAddr)
		ReloadEnv()
		assert.Equal(t, "Pouic", *addr)
		assert.Equal(t, "blou", *otherAddr)
	}
}

func TestGetenvInt(t *testing.T) {
	var addr *int
	var otherAddr *int
	// Test default value
	{
		t.Setenv("TEST_INT", "")
		addr = GetenvInt("TEST_INT", 42)
		assert.Equal(t, 42, *addr)
	}

	// Test if the address is the same
	{
		val := GetenvInt("TEST_INT", 100)
		assert.Equal(t, 42, *val)
		assert.Equal(t, addr, val)
	}

	// Test other env
	{
		t.Setenv("TEST_OTHER_INT", "123")
		otherAddr = GetenvInt("TEST_OTHER_INT", 42)
		assert.Equal(t, 123, *otherAddr)
		assert.Assert(t, addr != otherAddr)
	}

	// Test if the address is the same after reload
	{
		t.Setenv("TEST_INT", "99")
		t.Setenv("TEST_OTHER_INT", "456")
		assert.Equal(t, 42, *addr)
		assert.Equal(t, 123, *otherAddr)
		ReloadEnv()
		assert.Equal(t, 99, *addr)
		assert.Equal(t, 456, *otherAddr)
	}
}

func TestGetenvFloat64(t *testing.T) {
	var addr *float64
	var otherAddr *float64
	// Test default value
	{
		t.Setenv("TEST_FLOAT", "")
		addr = GetenvFloat64("TEST_FLOAT", 3.14)
		assert.Equal(t, 3.14, *addr)
	}

	// Test if the address is the same
	{
		val := GetenvFloat64("TEST_FLOAT", 1.23)
		assert.Equal(t, 3.14, *val)
		assert.Equal(t, addr, val)
	}

	// Test other env
	{
		t.Setenv("TEST_OTHER_FLOAT", "2.718")
		otherAddr = GetenvFloat64("TEST_OTHER_FLOAT", 3.14)
		assert.Equal(t, 2.718, *otherAddr)
		assert.Assert(t, addr != otherAddr)
	}

	// Test if the address is the same after reload
	{
		t.Setenv("TEST_FLOAT", "1.618")
		t.Setenv("TEST_OTHER_FLOAT", "0.577")
		assert.Equal(t, 3.14, *addr)
		assert.Equal(t, 2.718, *otherAddr)
		ReloadEnv()
		assert.Equal(t, 1.618, *addr)
		assert.Equal(t, 0.577, *otherAddr)
	}
}

func TestGetenvBool(t *testing.T) {
	var addr *bool
	var otherAddr *bool
	// Test default value
	{
		t.Setenv("TEST_BOOL", "")
		addr = GetenvBool("TEST_BOOL", true)
		assert.Equal(t, true, *addr)
	}

	// Test if the address is the same
	{
		val := GetenvBool("TEST_BOOL", false)
		assert.Equal(t, true, *val)
		assert.Equal(t, addr, val)
	}

	// Test other env
	{
		t.Setenv("TEST_OTHER_BOOL", "false")
		otherAddr = GetenvBool("TEST_OTHER_BOOL", true)
		assert.Equal(t, false, *otherAddr)
		assert.Assert(t, addr != otherAddr)
	}

	// Test if the address is the same after reload
	{
		t.Setenv("TEST_BOOL", "false")
		t.Setenv("TEST_OTHER_BOOL", "true")
		assert.Equal(t, true, *addr)
		assert.Equal(t, false, *otherAddr)
		ReloadEnv()
		assert.Equal(t, false, *addr)
		assert.Equal(t, true, *otherAddr)
	}
}

func TestGetenvStringSlice(t *testing.T) {
	var addr *[]string
	var otherAddr *[]string
	// Test default value
	{
		t.Setenv("TEST_STRING_SLICE", "")
		addr = GetenvStringSlice("TEST_STRING_SLICE", []string{"a", "b", "c"})
		assert.DeepEqual(t, []string{"a", "b", "c"}, *addr)
	}

	// Test if the address is the same
	{
		val := GetenvStringSlice("TEST_STRING_SLICE", []string{"x", "y", "z"})
		assert.DeepEqual(t, []string{"a", "b", "c"}, *val)
		assert.Equal(t, addr, val)
	}

	// Test other env
	{
		t.Setenv("TEST_OTHER_STRING_SLICE", "one,two,three")
		otherAddr = GetenvStringSlice("TEST_OTHER_STRING_SLICE", []string{"a", "b", "c"})
		assert.DeepEqual(t, []string{"one", "two", "three"}, *otherAddr)
		assert.Assert(t, addr != otherAddr)
	}

	// Test if the address is the same after reload
	{
		t.Setenv("TEST_STRING_SLICE", "x,y,z")
		t.Setenv("TEST_OTHER_STRING_SLICE", "1,2,3")
		assert.DeepEqual(t, []string{"a", "b", "c"}, *addr)
		assert.DeepEqual(t, []string{"one", "two", "three"}, *otherAddr)
		ReloadEnv()
		assert.DeepEqual(t, []string{"x", "y", "z"}, *addr)
		assert.DeepEqual(t, []string{"1", "2", "3"}, *otherAddr)
	}
}

func TestGetenvIntSlice(t *testing.T) {
	var addr *[]int
	var otherAddr *[]int
	// Test default value
	{
		t.Setenv("TEST_INT_SLICE", "")
		addr = GetenvIntSlice("TEST_INT_SLICE", []int{1, 2, 3})
		assert.DeepEqual(t, []int{1, 2, 3}, *addr)
	}

	// Test if the address is the same
	{
		val := GetenvIntSlice("TEST_INT_SLICE", []int{4, 5, 6})
		assert.DeepEqual(t, []int{1, 2, 3}, *val)
		assert.Equal(t, addr, val)
	}

	// Test other env
	{
		t.Setenv("TEST_OTHER_INT_SLICE", "10,20,30")
		otherAddr = GetenvIntSlice("TEST_OTHER_INT_SLICE", []int{1, 2, 3})
		assert.DeepEqual(t, []int{10, 20, 30}, *otherAddr)
		assert.Assert(t, addr != otherAddr)
	}

	// Test if the address is the same after reload
	{
		t.Setenv("TEST_INT_SLICE", "4,5,6")
		t.Setenv("TEST_OTHER_INT_SLICE", "40,50,60")
		assert.DeepEqual(t, []int{1, 2, 3}, *addr)
		assert.DeepEqual(t, []int{10, 20, 30}, *otherAddr)
		ReloadEnv()
		assert.DeepEqual(t, []int{4, 5, 6}, *addr)
		assert.DeepEqual(t, []int{40, 50, 60}, *otherAddr)
	}
}

func TestGetenvFloat64Slice(t *testing.T) {
	var addr *[]float64
	var otherAddr *[]float64
	// Test default value
	{
		t.Setenv("TEST_FLOAT_SLICE", "")
		addr = GetenvFloat64Slice("TEST_FLOAT_SLICE", []float64{1.1, 2.2, 3.3})
		assert.DeepEqual(t, []float64{1.1, 2.2, 3.3}, *addr)
	}

	// Test if the address is the same
	{
		val := GetenvFloat64Slice("TEST_FLOAT_SLICE", []float64{4.4, 5.5, 6.6})
		assert.DeepEqual(t, []float64{1.1, 2.2, 3.3}, *val)
		assert.Equal(t, addr, val)
	}

	// Test other env
	{
		t.Setenv("TEST_OTHER_FLOAT_SLICE", "10.1,20.2,30.3")
		otherAddr = GetenvFloat64Slice("TEST_OTHER_FLOAT_SLICE", []float64{1.1, 2.2, 3.3})
		assert.DeepEqual(t, []float64{10.1, 20.2, 30.3}, *otherAddr)
		assert.Assert(t, addr != otherAddr)
	}

	// Test if the address is the same after reload
	{
		t.Setenv("TEST_FLOAT_SLICE", "4.4,5.5,6.6")
		t.Setenv("TEST_OTHER_FLOAT_SLICE", "40.4,50.5,60.6")
		assert.DeepEqual(t, []float64{1.1, 2.2, 3.3}, *addr)
		assert.DeepEqual(t, []float64{10.1, 20.2, 30.3}, *otherAddr)
		ReloadEnv()
		assert.DeepEqual(t, []float64{4.4, 5.5, 6.6}, *addr)
		assert.DeepEqual(t, []float64{40.4, 50.5, 60.6}, *otherAddr)
	}
}

func TestGetenvBoolSlice(t *testing.T) {
	var addr *[]bool
	var otherAddr *[]bool
	// Test default value
	{
		t.Setenv("TEST_BOOL_SLICE", "")
		addr = GetenvBoolSlice("TEST_BOOL_SLICE", []bool{true, false, true})
		assert.DeepEqual(t, []bool{true, false, true}, *addr)
	}

	// Test if the address is the same
	{
		val := GetenvBoolSlice("TEST_BOOL_SLICE", []bool{false, true, false})
		assert.DeepEqual(t, []bool{true, false, true}, *val)
		assert.Equal(t, addr, val)
	}

	// Test other env
	{
		t.Setenv("TEST_OTHER_BOOL_SLICE", "false,true,false")
		otherAddr = GetenvBoolSlice("TEST_OTHER_BOOL_SLICE", []bool{true, false, true})
		assert.DeepEqual(t, []bool{false, true, false}, *otherAddr)
		assert.Assert(t, addr != otherAddr)
	}

	// Test if the address is the same after reload
	{
		t.Setenv("TEST_BOOL_SLICE", "false,true,false")
		t.Setenv("TEST_OTHER_BOOL_SLICE", "true,false,true")
		assert.DeepEqual(t, []bool{true, false, true}, *addr)
		assert.DeepEqual(t, []bool{false, true, false}, *otherAddr)
		ReloadEnv()
		assert.DeepEqual(t, []bool{false, true, false}, *addr)
		assert.DeepEqual(t, []bool{true, false, true}, *otherAddr)
	}
}

func TestRegisterCallbackEnv(t *testing.T) {
	callBackCalled := 0
	RegisterCallbackEnv("TEST_CALLBACK", func() {
		callBackCalled++
	})
	ReloadEnv()
	ReloadEnv()
	assert.Equal(t, 2, callBackCalled)
	UnregisterCallbackEnv("TEST_CALLBACK")
	ReloadEnv()
	assert.Equal(t, 2, callBackCalled)
	RegisterCallbackEnv("TEST_CALLBACK", func() {
		callBackCalled++
	})
	ReloadEnv()
	assert.Equal(t, 3, callBackCalled)
}
