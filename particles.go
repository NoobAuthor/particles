package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
	"time"
)

type Particle struct {
	x, y   float64
	vx, vy float64
	life   int
}

type ParticleSystem struct {
	particles []*Particle
	rnd       *rand.Rand
}

func (ps *ParticleSystem) Update() {
	for _, p := range ps.particles {
		p.x += p.vx
		p.y += p.vy
		p.vx *= 0.99
		p.life -= 1

		if p.life <= 0 {
			p.x = 320
			p.y = 240
			p.vx = rand.Float64()*2 - 1
			p.vy = rand.Float64()*2 - 1
			p.life = rand.Intn(60) + 60
		}
	}
}

func (ps *ParticleSystem) Draw(screen *ebiten.Image) {
	rectImg := ebiten.NewImage(2, 2)
	rectImg.Fill(color.White)

	for _, p := range ps.particles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(p.x, p.y)
		screen.DrawImage(rectImg, op)
	}
}

func (ps *ParticleSystem) createParticle() {
	p := &Particle{
		x:    320,
		y:    240,
		vx:   rand.Float64()*2 - 1,
		vy:   rand.Float64()*2 - 1,
		life: rand.Intn(60) + 60,
	}
	ps.particles = append(ps.particles, p)
}

type Game struct {
	ps ParticleSystem
}

func (g *Game) Update() error {
	g.ps.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ps.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	game := &Game{
		ps: ParticleSystem{
			rnd: rnd,
		},
	}

	for i := 0; i < 100; i++ {
		game.ps.createParticle()
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Particle Simulator")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
