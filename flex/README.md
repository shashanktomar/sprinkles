# Flex

TODO GIF

Check [a guide to flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox/)
to understand the details.

Currently support:

- row and column directions on a container
- `grow`, `shrink` and `basis` properties on a box
- `minsize` and `maxsize` on a box

## Examples

You can run examples under `./examples` like `cd ./examples/flex-row/ && go run .`

## Tutorial

Create a container and add boxes:

```go
layout := flex.NewContainer(flex.Row).
  AddBox(boxOne, flex.NewStyle().Flex(1)).
  AddBox(boxTwo, flex.NewStyle().Flex(2)).
  AddBox(boxThree, flex.NewStyle().Flex(4))
```

Here a box need to implement:

```go
type Box interface {
 SetSize(int, int)
 View() string
}
```

### Grow in row direction

### Grow in column direction

### Composed in both directions

## Potential Roadmap

Container

- [x] add row and column direction properties
- [ ] add flex-wrap property
- [ ] add justify-content property
- [ ] add align-items property
- [ ] add align-content property
- [ ] add gap, row-gap and column-gap properties
- [ ] add row-reverse and column-reverse properties

Box

- [x] add flex grow property
- [x] add flex shrink property
- [x] add flex basis property
- [ ] add order property
- [ ] add alignment property
