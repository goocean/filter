
# filter


### ParamsFromValues(v url.Values, b bool)

  Filter Params form Values for `r.URL.Query()` in [pat][], [gohttp][], [bone][].
  If b = true, delete Params from Values.

```go
v := url.Values{}
v.Set(":name", "Ava")
v.Set("name", "Zoe")
b := false
p := ParamsFromValues(v, b)
// p.Get("name")
```


[pat]: https://github.com/bmizerany/pat
[gohttp]: https://github.com/gohttp
[bone]: https://github.com/squiidz/bone