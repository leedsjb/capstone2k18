import { injectGlobal } from "styled-components";

/* eslint no-unused-expressions: 0 */
injectGlobal`
    body {
        font-family: sans-serif;
    }

    #root {
        display: flex;
        flex-direction: column;
    }

    .active {}

    a, a:visited, a:hover, a:active {
        color: inherit;
        text-decoration: none;
    }
`;
