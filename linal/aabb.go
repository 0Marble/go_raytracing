package linal

type Aabb struct {
	Min Vec3
	Max Vec3
}

func (a *Aabb) ContainsPoint(pt Vec3) bool {
	return pt.X >= a.Min.X && pt.X <= a.Max.X &&
		pt.Y >= a.Min.Y && pt.Y <= a.Max.Y &&
		pt.Z >= a.Min.Z && pt.Z <= a.Max.Z
}

func (a *Aabb) Merge(b *Aabb) Aabb {
	return Aabb{a.Min.Min(b.Min), a.Max.Max(b.Max)}
}
