package sprite

import "github.com/hajimehoshi/ebiten"

func isOverlap(x1, x2, x3, x4 int) bool {
	if x1 <= x4 && x2 >= x3 {
		return true
	}
	return false
}

type Player struct {
	BaseSprite
}

func NewPlayer(images []*ebiten.Image) *Player {
	player := new(Player)
	player.Images = images
	player.ImageNum = len(images)
	return player
}

func (p *Player) Move(objects []Sprite) {
	var dx, dy int
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dx = -1
		p.count++
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dx = 1
		p.count++
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		dy = -1
		p.count++
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		dy = 1
		p.count++
	}

	for _, object := range objects {
		dx, dy = p.IsCollide(dx, dy, object)
	}

	p.Position.X += dx
	p.Position.Y += dy
}

func (p *Player) IsCollide(dx, dy int, object Sprite) (int, int) {
	x := p.Position.X
	y := p.Position.Y
	img := p.currentImage()
	w, h := img.Size()

	x1, y1, w1, h1 := object.GetCordinates()

	overlappedX := isOverlap(x, x+w, x1, x1+w1)
	overlappedY := isOverlap(y, y+h, y1, y1+h1)

	if overlappedY {
		if dx < 0 && x+dx <= x1+w1 && x+w+dx >= x1 {
			dx = 0
		} else if dx > 0 && x+w+dx >= x1 && x+dx <= x1+w1 {
			dx = 0
		}
	}
	if overlappedX {
		if dy < 0 && y+dy <= y1+h1 && y+h+dy >= y1 {
			dy = 0
		} else if dy > 0 && y+h+dy >= y1 && y+dy <= y1+h1 {
			dy = 0
		}
	}

	return dx, dy
}
