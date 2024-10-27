# NOSTALGIA

A tribute to the Dark Priests

## gallery

screenshot #1
<p align="center">
  <img style="float: right;" src="https://github.com/fred1268/assets/blob/development/nostalgia/screenshot01.png" alt="Nostalgia #1"/>
</p>

screenshot #2
<p align="center">
  <img style="float: right;" src="https://github.com/fred1268/assets/blob/development/nostalgia/screenshot02.png" alt="Nostalgia #2"/>
</p>

## ebiwhat?

Recently, I stumble upon [Ebitengine](https://ebitengine.org/) which markets itself as "*A dead simple 2D game
engine for Go*". Well, since this is exactly the kind of things I may be using in a future project, I decided
to give it a try. Reading docs and looking at examples is nice, but it is far from being enough to have a good
grasp on something. Thus I decided to write a 1980s-like megademo using Ebitengine.

## megawatt?

Back in the 80s, it was common to see 4k and 64k demos as well as megademos. As far as I remember,
[The demoscene](https://www.demoscene.info/) started on C64, then later migrated to Amiga and Atari, to end
up on PC.

More than 30 years ago, after having already written literally hundreds of thousands of lines of Assembly, C and
Pascal, and having been drooling over these demos on C64, Amiga and Atari, I convinced a couple of friends to create
what was probably the first French demo group on PC, and to jail ourselves in my home for a good long weekend
to create our first (and last) 1.44MiB megademo.

There was already some French groups on Amiga (hello [Wild Copper Crew](https://www.youtube.com/watch?v=OmhR7q2BwSY&t=78s) :wave:)
and, we discoved them later, a very talented Finish group on PC (hello [Future Crew](https://www.youtube.com/watch?v=InrGJ7C9B3s) :wave:),
but this did not deter us from our goal.

## a little bit of history (for those who care)

This was the time of the 386 and the unfamous 640k RAM limit (plus potential XMS), a time where every CPU
cycle and every bit of memory was the most precious thing. To give you an idea of both, here are two examples:

- When one needed to put zero in a register, the most obvious thing was to write something like `mov ax,0`,
which was translated into something like `B8 00 00`, `B8` to move a 16bits immediate value into `ax` and `00 00`
being the 16bits representation of 0. That's 3 bytes! 3 fat bytes! But if you write something like `xor ax,ax` or
`sub ax,ax` you endup with a single byte of code: a 66% saving (not counting the CPU cycle saving)!
Damn, you feel good when you discover this!

- Another example was related to writing to the video buffer. At that time, the MCGA mode (320x200, 256 paletted
colors) was the norm, and to light the pixel at (x,y) with color c, you had to write c at 0xA000+320*y+x. Easy
enough, except... that multiplication was not an option (if I recall well a `mul` was costing about 80 CPU cycles...).
So you had to be clever, and realize that 320=256+64. Thus, you could write this multiplication like this
(assuming y is in ax and x in bx):

```asm
    mov     dx, ax  // copy ax (which contains y) to dx
    shl     ax, 8   // multiply ax by 256 by shifting bits 8 times left
    shl     dx, 6   // multiply dx by 64 by shifting bits 6 times left
    add     ax, dx  // add both
    add     ax, bx  // add bx (which contains x)
```

...and you have the result of y*320+x for a few CPU cycles only! Magical! This also makes you feel good when
you realize it!

## how megademos were built?

Megademos have obviously evolved with time, but there were some things that you would find in pretty much every
single demo: star fields, scroll texts, sinusoidal scroll texts, sprites, and a little bit of 3D since we were at
the very beginning of it. And without GPU of course (no bitcoin mining back then!).

To achieve this, all the demos were written in Assembly, with the tooling made in C and / or Pascal. By tooling, I
mean that some of what was rendered on screen was precomputed. For instance, it was just not possible to compute
a sinus in a decent time, nor to make floating point calculations (the *mathematical co-processor* as it was
called, the [*i387*](https://en.wikipedia.org/wiki/X87#80387), was a very costly option, rarely found in
computers). So to solve this, we were doing fixed point arthmetics with precalculated sinus tables.

For instance, a C or Pascal program would be created to generate a Assembly file containing sinus values from 0 to π
by the increment of your choice (the greater the array, the higher the resolution), and multiplied by 256 to
keep each value in a single byte. When we had room, we could use 16bits amplitude rather than 8bits for a more
accurate result.

Then the Assembly program was making the real time part of the computation with integers, calculating the position
of each stripe of the font or each sprite, while handling vertical sync, music, etc.

## back to Nostalgia

My goal with Nostalgia was to pay tribute to my old friends, MagicManu++, ChD and Xcalibur, by trying
to reproduce some of these effects with Ebitengine, the way we did them at that time, while, at the same
time learning to use *A dead simple 2D game engine for Go*.

I had to cheat a little bit by using matrix because that's the way Ebitengine works, but I refrained from using
Shaders or more advanced 3D like triangles or vertices, that are also available in Ebitengine. As you can guess,
I also used real time math functions and extensively used `float64`... why wouldn't I?

So do not expect a state of the art use of 3D, this was not the point this time. I may try some more advanced
features of Ebitengine in the future.

## tweaking

If you would like to turn some knobs and tweak some parameters, most of the configuration is done in the
`internal/demo` package. I cannot promise that you won't run into an error, since I did not make my code
100% bulletproof (again, that was not the goal of the exercise), but you should be able, with a little
bit of understanding of the code in the `internal/text` and `internal/gfx` modules, to adapt this
microdemo to your likings.

Enjoy!

## additional credits

- The font comes from [TextCraft](https://textcraft.net/).
- The music comes from [Of Far Different Nature](https://fardifferent.itch.io/loops).
- And... I designed the sprites myself using [Gimp](https://www.gimp.org/): that's, in itself, quite an achievement :joy:!

## (weekend) conclusion

I will confess, it was quite a blast to code this microdemo in a weekend!

:wave: Hail MagicManu++, ChD and Xcalibur my friends: This is for you!
