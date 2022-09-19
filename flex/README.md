# Flex

Check [a guide to flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox/) to understand the details.

## Properties

Container

- `direction`: either `row` or `column`. The direction of the flex container. See examples for more details

Box

- `ratio`: a positive number. The ratio of the box within the container, relative to other boxes.
  If there are three boxes with ratio 1, 1, 3. The first two boxes will have same size and the third one will be 3 times the size.

## Potential Roadmap

Container

- [x] add direction property
- [ ] add flex-wrap property
- [ ] add justify-content property
- [ ] add align-items property
- [ ] add align-content property
- [ ] add gap, row-gap and column-gap properties

Box

- [x] add flex ratio property
- [ ] add order property
- [ ] add flex shrink property
- [ ] add flex basis property
- [ ] add align self property
