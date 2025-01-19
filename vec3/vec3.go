package vec3

import "sknoslo/ebc2024/math"

type Vec3 struct {
	X, Y, Z int
}

func New(x, y, z int) Vec3 {
	return Vec3{x, y, z}
}

func (v Vec3) InRange(x1, y1, x2, y2, z1, z2 int) bool {
	return v.X >= x1 && v.X <= x2 && v.Y >= y1 && v.Y <= y2 && v.Z >= z1 && v.Z <= z2
}

func (va Vec3) Add(vb Vec3) Vec3 {
	return Vec3{va.X + vb.X, va.Y + vb.Y, va.Z + vb.Z}
}

func (va Vec3) Sub(vb Vec3) Vec3 {
	return Vec3{va.X - vb.X, va.Y - vb.Y, va.Z - vb.Z}
}

func (v Vec3) Mul(s int) Vec3 {
	return Vec3{v.X * s, v.Y * s, v.Z * s}
}

func (v Vec3) Div(s int) Vec3 {
	return Vec3{v.X / s, v.Y / s, v.Z / s}
}

func Distance(a, b Vec3) int {
	return math.AbsDiff(a.X, b.X) + math.AbsDiff(a.Y, b.Y) + math.AbsDiff(a.Z, b.Z)
}
