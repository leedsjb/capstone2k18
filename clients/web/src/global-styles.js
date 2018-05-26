import { injectGlobal } from "styled-components";

/* eslint no-unused-expressions: 0 */
injectGlobal`
    body {
        font-family: sans-serif;
        line-height: 1.5;
    }

    h1, h2, h3, h4, h5, h6 {
        line-height: 1.25;
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

    .mapboxgl-popup-content {
        padding: 4px 12px;
    }

    
`;
