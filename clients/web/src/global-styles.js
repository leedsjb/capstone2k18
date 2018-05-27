import { injectGlobal } from "styled-components";

/* eslint no-unused-expressions: 0 */
injectGlobal`
    * {
        border: 0;
        box-sizing: inherit;
        -webkit-font-smoothing: antialiased;
        font-weight: inherit;
        margin: 0;
        outline: 0;
        padding: 0;
        text-decoration: none;
        text-rendering: optimizeLegibility;
        -webkit-appearance: none;
        -moz-appearance: none;
        min-height: 0;
    }
    
    html {
        display: flex;
        height: 100%;
        width: 100%;
        max-height: 100%;
        max-width: 100%;
        box-sizing: border-box;
        font-size: 16px;
        line-height: 1.5;
        background-color: #ffffff;
        padding: 0;
        margin: 0;
        -webkit-font-smoothing: antialiased;
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
    }

    body {
        display: flex;
        box-sizing: border-box;
        flex: auto;
        align-self: stretch;
        max-width: 100%;
        max-height: 100%;
        -webkit-overflow-scrolling: touch;
    }

    #root {
        display: flex;
        display: -webkit-box;
        display: -webkit-flex;
        display: -moz-flex;
        display: -ms-flexbox;
        flex-direction: column;
        -ms-flex-direction: column;
        -moz-flex-direction: column;
        -webkit-flex-direction: column;
        height: 100%;
        width: 100%;
    }

    h1, h2, h3, h4, h5, h6 {
        line-height: 1.25;
    }

    a, a:visited, a:hover, a:active {
        color: inherit;
        text-decoration: none;
    }

    .mapboxgl-popup-content {
        padding: 4px 12px;
    }
`;
