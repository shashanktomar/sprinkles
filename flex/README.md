# Flex

![master-gif](https://user-images.githubusercontent.com/2309241/193480567-ec905b89-7c87-4a55-b6bb-cca2ad1c8969.gif)

Check [a guide to flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox/)
to understand the details.

Currently support:

- row and column directions on a container
- `grow`, `shrink` and `basis` properties on a box
- `minsize` and `maxsize` on a box

## Examples

You can run examples under `./examples` like `cd ./examples/flex-row/ && go run .`

## Tutorial

### Setup

Create a container and add boxes:

```go
layout := flex.NewContainer(flex.Row).
  AddBox(boxOne, flex.NewStyle().Flex(1)).
  AddBox(boxTwo, flex.NewStyle().Flex(2)).
  AddBox(boxThree, flex.NewStyle().Flex(4))
```

The view you are putting in the container need to implement Box interface:

```go
type Box interface {
 SetSize(int, int)
 View() string
}
```

### Invoke SetSize on Container

Call `layout.SetSize(w, h)` when you receive an `Update` with `tea.WindowSizeMsg` message.
This triggers the flex calculation on all children on the container and they are assigned
appropriate size by invoking `box.SetSize(w,h)` method. This is where you should assign size
to your view.

### Invoke View on Container

Call `layout.View()` to render all the boxes.

## How does it work?

A **Container** is created with a direction. This decides the direction in which flex behaviour
will be applied:
- `Direction`: `row` or `column`

A **Box** has following properties:
- `grow`: This influence the proportion in which the box will grow as compared to other boxes. For
example, if there are three boxes with grow as 1,6,3 and there are 100 units of free space available,
they will be assigned 10, 60 and 30 units respectively.
- `basis`: the base size of the container. When flex behaviour is applied, the calculation for a box start
with `basis` and then free space (or negative space) is applied on top of the `basis` based on `grow` or `shrink`
properties of the box.
- `shrink`: This influence the proportion in which the box will shrink as compared to other boxes. The
shrink only happen if there is negative space i.e. the total `basis` of all the boxes is bigger than assigned
width on the container.
- `maxSize`: a container can not grow beyond maxSize
- `minSize`: a container can not shrink beyond minSize

## Potential Roadmap

Container

- [x] add row and column direction properties
- [ ] support style on container
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
