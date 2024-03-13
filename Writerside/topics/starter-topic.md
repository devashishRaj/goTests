# CH1. Tools for Testing

>When I write a test, half that test is not just to validate the code is working as expected,
> it’s to validate the API is intuitive and simple to use.   
> —Bill Kennedy, “Package Oriented Design”

#### Concurrent tests with t.Parallel()
- any test can declare itself suitable to run concurrently with other tests,
you can speed up your tests by running them in parallel like this
```Go
func TestCouldTakeAWhile(t *testing.T) {
    t.Parallel()
    ... // this could take a while
}
```

#### Failures: t.Error and t.Errorf 
> they both mark marks the test as failed, but does not stop it so the test will continue  
> A good test should not give up after one failure but should try to report several errors in a single run, 
> since the pattern of failures may itself be revealing.
- many Go functions return error along with some data value, it’s as important to test that error result as any other.
```Go
if err != nil {
    t.Error(err)
}
```
OR
```Go
want := 4
got := Add(2, 2)
if want != got {
    t.Error("want", want, "got", got)
}
```
- If we need more precise control over the failure data, then we can use t.Errorf instead.
```Go
want := "hello"
got := Greeting()
if want != got {
    t.Errorf("want %q, got %q", want, got)
}
```

#### Abandoning the test with t.Fatal
- use fatal when continuing does not matter say, you need to access some file first before checking behaviour of some 
function , is OS is failing to access the file. There's no point to continue the test.
```Go
f, err := os.Open("testdata/input")
if err != nil {
    t.Fatal(err)
}
```

#### Writing debug output with t.Log
- sometimes we’d like the test to be able to output things that aren’t failure messages
```Go
// to know intermediate values
// if the test passes , we will see no output but if it fails, all calls to t.Log() will be shown
// as tests should be silent on passing and loud on failing
got := StageOne()
t.Log("StageOne result", got)
got = StageTwo(got)
t.Log("StageTwo result", got)
got = StageThree(got)
if want != got {
    t.Errorf("want %v, got %v", want, got)
} 
```

#### Test flags: -v and -run





