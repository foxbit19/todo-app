# Easy to do - client

This client is implemented in React.

The Easy to do offers a simple way to manage a list of tasks someone needs to complete or things that someone wants to do.

## Introduction
This react application is very simple. It contains the less code to work properly and to satisfy the requirement.

Despite that, I wanted to develop it well using `react 18`, `typescript` and 2 strong libraries that I'd love:

- `material-ui`, and
- `react-beautiful-dnd`.

The first, provided by the package `@mui/material`, was a huge collection of nice looking components ready to go.

The second is a library to provide drag & drop (dnd) operation on graphical component.
With this library I've encountered some issue (because it's not maintained well and I've used React 18) so I've integrated it with the package `@hello-pangea/dnd` that provides support for react 18.

I hope this application enjoy you, as I've enjoyed myself during its development.
## Structure
`src` is the main folder for source code and test files. All the test files were written in the same directory of the components.

- `src/components` contains all the component of the application, some were written for maximum readability (like `GenericDialog`), others for code separation and maintainability;
- `src/models` contains the model `Item` used by the application to store the field of an item and the adapter `ItemAdapter` to adapt json results from the server to the local model;
- `src/services` contains all the code necessary to interact with the server:
  - `service.ts` contains an interface that works as a guide for all service implementations. I've used generics to develop this one: a service that implements `Service` needs to specify a type to map argument and returns of its typed functions.
  - `itemService.ts` contains the class `ItemService` that implements `Service` with type `Item`
  - `completedItemService.ts` contains the class `CompletedItemService` that implements `Service` with type `Item`. Even if this interface does not use different typed argument respect `ItemService` it was created to show the potential of a different implementation and to separate concern. It works only with completed items. It could be extended in the future to interact only with this kind of elements avoiding to create out-of-scope function inside `ItemService`;

## Usage

### Install
Install the dependencies using the command

```bash
npm install
```

### Start the application

Start the application using the command

```bash
npm start
```

#### Environment
As for `.env` files it uses `http://localhost:8080` as base url and it exposes port `3000` by default.

### Testing

All the test use `jest` for component testing. It's a way to test the behaviour of isolated components but it's not an end-to-end test.

#### How to test
Launch the test using the command

```bash
npm run test
```

## Improvements
There are one big improvement to this application: end-to-end tests.

It is possible to integrate these tests using framework like `cypress.js`. I've made some experiments but I've encountered some issues caused by `tsconfig.json` configuration.