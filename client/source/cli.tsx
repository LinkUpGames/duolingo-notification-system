#!/usr/bin/env node
import React from "react";
import meow from "meow";
import App from "./app.js";
import { withFullScreen } from "fullscreen-ink";

const cli = meow(
  `
	Usage
	  $ client

	Options
		--name  Your name

	Examples
	  $ client --name=Jane
	  Hello, Jane
`,
  {
    importMeta: import.meta,
    flags: {
      name: {
        type: "string",
      },
    },
  }
);

withFullScreen(<App name={cli.flags.name} />).start();
