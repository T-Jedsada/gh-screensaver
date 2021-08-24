# gh-screensaver

_being a gh extension that runs animated terminal "screensavers"_

## usage

- `gh screensaver` run a random screensaver
- `gh screensaver -s pipes` run a screensaver by name
- `gh screensaver -l` list available screensavers

Extra configuration options can be passed after a `--`; for example:

```
gh screensaver -smarquee -- --message="hello world" --font="script"
```

## savers

### fireworks

watch a fireworks display.

![fwork2](https://user-images.githubusercontent.com/98482/130655005-f68305f9-e8c1-45c9-b998-de5167422136.gif)


`--color` `full` or `off`. Default `full`

### starfield

fly through space.

![starfield](https://user-images.githubusercontent.com/98482/130655039-7d76e84d-2eae-4347-a1fd-6c46e329bfda.gif)

`--density` Default `250`. The number of stars to render.
`--speed` Default `4`. Higher is faster.

### pipes

2d pipes draw across the screen.

![pipes2](https://user-images.githubusercontent.com/98482/130655085-bb797fc8-0c2a-4fda-926b-963d84a2d611.gif)


`--color` `full` or `off`. Default `full`

### pollock

paint splotches cover the screen.

![pollock](https://user-images.githubusercontent.com/98482/130655110-8704912e-936b-4a23-9bff-8eb1045a2488.gif)


## author

nate smith <vilmibm@github.com>
