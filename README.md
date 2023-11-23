# Go Spherical

[![Test](https://github.com/stellarsunset/go-spherical/actions/workflows/test.yaml/badge.svg)](https://github.com/stellarsunset/go-spherical/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/stellarsunset/go-spherical/graph/badge.svg?token=CW10XNEJA5)](https://codecov.io/gh/stellarsunset/go-spherical)

Math on a spherical Earth.

Quickly compute distances, courses, projections, etc. using locations on the Earth's surface.

```golang
import(
    sph "stellarsunset/spherical"
    ll "stellarsunset/spherical/latlong"
    dist "stellarsunset/spherical/distance"
)

// Basic operations against primitives, compute courses, distances, perform projections, etc.
nyc := []float64{40.7128, -74.0060}
tokyo := []float64{35.6764, 139.6500}

nm := sph.DistanceInNm(nyc[0], nyc[1], tokyo[0], tokyo[1])
fmt.Println(nm) // 5856.184170028993

deg := sph.CourseInDegrees(nyc[0], nyc[1], tokyo[0], tokyo[1])
fmt.Println(deg) // 332.98810818766503

// Operations using higher-level, but simple abstractions
nyc, tokyo := ll.NewLatLong(40.7128, -74.0060), ll.NewLatLong(35.6764, 139.6500)

d := nyc.DistanceTo(tokyo)
fmt.Println(d.InNauticalMiles()) // 5856.184170028993
fmt.Println(d.InFeet()) // 3.558285132182971e+07

c := nyc.CourseTo(tokyo)
fmt.Println(c.InDegrees()) // 332.98810818766503
fmt.Println(c.InRadians()) // 5.811738857861843
```

Much of the functionality in this repository has been ported from [MITRE Commons](https://github.com/mitre-public/commons).
