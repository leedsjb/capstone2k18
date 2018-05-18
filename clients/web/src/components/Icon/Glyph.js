import React from "react";

const Glyph = ({ glyph }) => {
    /* eslint-disable default-case */
    switch (glyph) {
        case "medicalCrossLine":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g>
                        <g
                            stroke="currentColor"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                        >
                            <polygon points="23.209 8.414 20.709 4.084 14.5 7.669 14.5 0.5 9.5 0.5 9.5 7.668 3.293 4.084 0.793 8.414 7 12 0.793 15.584 3.293 19.914 9.5 16.329 9.5 23.5 14.5 23.5 14.5 16.329 20.709 19.914 23.209 15.584 17 12" />
                        </g>
                    </g>
                </g>
            );
        case "medicalCrossFilled":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g>
                        <g fill="currentColor" fillRule="nonzero">
                            <path d="M23.459,15.15 L18,12 L23.459,8.847 C23.698,8.71 23.78,8.405 23.642,8.165 L21.142,3.834 C21.076,3.719 20.966,3.635 20.838,3.602 C20.71,3.567 20.573,3.584 20.459,3.651 L15,6.803 L15,0.5 C15,0.224 14.775,0 14.5,0 L9.5,0 C9.224,0 9,0.224 9,0.5 L9,6.801 L3.542,3.65 C3.427,3.584 3.29,3.566 3.163,3.602 C3.035,3.635 2.926,3.719 2.859,3.834 L0.359,8.165 C0.221,8.404 0.303,8.71 0.542,8.847 L6,12 L0.542,15.15 C0.303,15.29 0.221,15.594 0.359,15.834 L2.859,20.165 C2.925,20.278 3.035,20.362 3.163,20.396 C3.29,20.429 3.427,20.413 3.542,20.346 L9,17.196 L9,23.5 C9,23.776 9.224,24 9.5,24 L14.5,24 C14.775,24 15,23.776 15,23.5 L15,17.196 L20.459,20.346 C20.573,20.412 20.71,20.429 20.838,20.396 C20.966,20.362 21.075,20.278 21.142,20.165 L23.642,15.834 C23.708,15.719 23.726,15.582 23.692,15.454 C23.657,15.327 23.573,15.217 23.459,15.15 Z" />
                        </g>
                    </g>
                </g>
            );
        case "airplaneFlightLine":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g>
                        <g
                            transform="translate(0, 4)"
                            stroke="currentColor"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                        >
                            <path d="M21.5,5.5 L6.5,5.5 L4.5,3.5 L0.5,3.5 L3.5,9.5 L11.5,9.5 L5.5,15.5 L10.5,15.5 L16.5,9.5 L21.5,9.5 C22.604,9.5 23.5,8.604 23.5,7.5 C23.5,6.396 22.604,5.5 21.5,5.5 Z" />
                            <polyline points="16.5 5.5 11.5 0.5 7.5 0.5 12.5 5.5" />
                        </g>
                    </g>
                </g>
            );
        case "airplaneFlightFilled":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g>
                        <g
                            transform="translate(0, 4)"
                            fill="currentColor"
                            stroke="currentColor"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                        >
                            <path d="M21.5,5.5 L6.5,5.5 L4.5,3.5 L0.5,3.5 L3.5,9.5 L11.5,9.5 L5.5,15.5 L10.5,15.5 L16.5,9.5 L21.5,9.5 C22.604,9.5 23.5,8.604 23.5,7.5 C23.5,6.396 22.604,5.5 21.5,5.5 Z" />
                            <polyline points="16.5 5.5 11.5 0.5 7.5 0.5 12.5 5.5" />
                        </g>
                    </g>
                </g>
            );
        case "accountGroupLine":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g>
                        <g
                            transform="translate(0, 1)"
                            stroke="currentColor"
                            strokeLinejoin="round"
                        >
                            <path d="M0.5,7.5 L0.5,10.5 C0.5,11.325 0.675,12 1.5,12 L1.5,16.5 L4.5,16.5 L4.5,12 C5.324,12 5.499,11.324 5.5,10.5 L5.5,7.5 L0.5,7.5 Z" />
                            <path d="M18.5,7.5 L18.5,10.5 C18.5,11.325 18.675,12 19.5,12 L19.5,16.5 L22.5,16.5 L22.5,12 C23.324,12 23.499,11.324 23.5,10.5 L23.5,7.5 L18.5,7.5 Z" />
                            <path d="M8.5,7.5 L15.5,7.5 L15.5,12.5 C15.5,13.6 14.6,14.5 13.5,14.5 L13.5,20.5 L10.5,20.5 L10.5,14.5 C9.399,14.5 8.5,13.6 8.5,12.5 L8.5,7.5 Z" />
                            <path d="M4.5,4 C4.5,4.828 3.828,5.5 3,5.5 C2.172,5.5 1.5,4.828 1.5,4 C1.5,3.172 2.172,2.5 3,2.5 C3.828,2.5 4.5,3.172 4.5,4 Z" />
                            <path d="M22.5,4 C22.5,4.828 21.828,5.5 21,5.5 C20.172,5.5 19.5,4.828 19.5,4 C19.5,3.172 20.172,2.5 21,2.5 C21.828,2.5 22.5,3.172 22.5,4 Z" />
                            <path d="M14.5,3 C14.5,4.381 13.381,5.5 12,5.5 C10.619,5.5 9.5,4.381 9.5,3 C9.5,1.619 10.619,0.5 12,0.5 C13.381,0.5 14.5,1.619 14.5,3 Z" />
                        </g>
                    </g>
                </g>
            );
        case "close":
            return (
                <g stroke="none" strokeWidth="1" fillRule="evenodd">
                    <path d="M13.789147,12 L23.6294555,2.15969148 C24.1235148,1.66563218 24.1235148,0.86460378 23.6294555,0.370544477 C23.1353962,-0.123514826 22.3343678,-0.123514826 21.8403085,0.370544477 L12,10.210853 L2.15969148,0.370544477 C1.66563218,-0.123514826 0.86460378,-0.123514826 0.370544477,0.370544477 C-0.123514826,0.86460378 -0.123514826,1.66563218 0.370544477,2.15969148 L10.210853,12 L0.370544477,21.8403085 C-0.123514826,22.3343678 -0.123514826,23.1353962 0.370544477,23.6294555 C0.86460378,24.1235148 1.66563218,24.1235148 2.15969148,23.6294555 L12,13.789147 L21.8403085,23.6294555 C22.3343678,24.1235148 23.1353962,24.1235148 23.6294555,23.6294555 C24.1235148,23.1353962 24.1235148,22.3343678 23.6294555,21.8403085 L13.789147,12 Z" />
                </g>
            );
        case "accountGroupFilled":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g>
                        <g
                            transform="translate(0, 1)"
                            fill="currentColor"
                            fillRule="nonzero"
                        >
                            <path d="M5.5,7 L0.5,7 C0.224,7 0,7.224 0,7.5 L0,10.5 C0,11.583 0.336,12.23 1,12.432 L1,16.5 C1,16.776 1.224,17 1.5,17 L4.5,17 C4.776,17 5,16.776 5,16.5 L5,12.432 C5.663,12.23 5.998,11.583 6,10.5 L6,7.5 C6,7.224 5.776,7 5.5,7 Z" />
                            <path d="M23.5,7 L18.5,7 C18.224,7 18,7.224 18,7.5 L18,10.5 C18,11.583 18.336,12.23 19,12.432 L19,16.5 C19,16.776 19.224,17 19.5,17 L22.5,17 C22.776,17 23,16.776 23,16.5 L23,12.432 C23.663,12.23 23.999,11.583 24,10.5 L24,7.5 C24,7.224 23.776,7 23.5,7 Z" />
                            <path d="M15.5,7 L8.5,7 C8.224,7 8,7.224 8,7.5 L8,12.5 C8,13.708 8.86,14.717 10,14.95 L10,20.5 C10,20.776 10.224,21 10.5,21 L13.5,21 C13.776,21 14,20.776 14,20.5 L14,14.95 C15.14,14.718 16,13.708 16,12.5 L16,7.5 C16,7.224 15.776,7 15.5,7 Z" />
                            <circle cx="3" cy="4" r="2" />
                            <circle cx="21" cy="4" r="2" />
                            <circle cx="12" cy="3" r="3" />
                        </g>
                    </g>
                </g>
            );
        case "navigationDrawerFilled":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g>
                        <g
                            transform="translate(4, 5)"
                            fill="currentColor"
                            fillRule="nonzero"
                        >
                            <path d="M1.043,2 L16,2 C16.553,2 17,1.552 17,1 C17,0.448 16.553,0 16,0 L1.043,0 C0.49,0 0.043,0.448 0.043,1 C0.043,1.552 0.49,2 1.043,2 Z" />
                            <path d="M16,4 L1.043,4 C0.49,4 0.043,4.448 0.043,5 C0.043,5.552 0.49,6 1.043,6 L16,6 C16.553,6 17,5.552 17,5 C17,4.448 16.553,4 16,4 Z" />
                            <path d="M16,8 L1.043,8 C0.49,8 0.043,8.448 0.043,9 C0.043,9.552 0.49,10 1.043,10 L16,10 C16.553,10 17,9.552 17,9 C17,8.448 16.553,8 16,8 Z" />
                            <path d="M16,12 L1.043,12 C0.49,12 0.043,12.448 0.043,13 C0.043,13.552 0.49,14 1.043,14 L16,14 C16.553,14 17,13.552 17,13 C17,12.448 16.553,12 16,12 Z" />
                        </g>
                    </g>
                </g>
            );
        case "chevronLeft":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g>
                        <g transform="translate(4, 0)" fill="currentColor">
                            <path d="M5.91181176,11.52192 L15.3715765,3.32832 C15.4666353,3.24576 15.5259294,3.12768 15.5362824,3 C15.5456941,2.87328 15.5061647,2.74656 15.4242824,2.65056 L13.3414588,0.17184 C13.2605176,0.07392 13.1438118,0.01536 13.0195765,0.0048 C12.8981647,-0.00384 12.7711059,0.03744 12.6760471,0.12096 L0.170635294,11.15904 C0.0671058824,11.25024 0.00781176471,11.38176 0.00781176471,11.52192 C0.00781176471,11.66208 0.0671058824,11.7936 0.170635294,11.8848 L12.6760471,22.92672 C12.7626353,23.00256 12.8718118,23.04384 12.9847529,23.04384 C12.9960471,23.04384 13.0073412,23.04384 13.0195765,23.04288 C13.1438118,23.03232 13.2605176,22.97376 13.3414588,22.87584 L15.4242824,20.39712 C15.5061647,20.30112 15.5456941,20.1744 15.5362824,20.04768 C15.5259294,19.92 15.4675765,19.80192 15.3715765,19.71936 L5.91181176,11.52192 Z" />
                        </g>
                    </g>
                </g>
            );
        case "map":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g>
                        <g
                            transform="translate(2, 1)"
                            fill="currentColor"
                            fillRule="nonzero"
                        >
                            <path d="M0.27,0.061 C0.105,0.148 0,0.319 0,0.506 L0,16.506 C0,16.668 0.078,16.819 0.209,16.913 L7,21.764 L7,4.535 L0.79,0.099 C0.638,-0.009 0.438,-0.025 0.27,0.061 Z" />
                            <path d="M21.79,5.099 L15,0.25 L15,17.479 L21.209,21.914 C21.359,22.021 21.558,22.039 21.729,21.952 C21.896,21.866 22,21.693 22,21.506 L22,5.506 C22,5.345 21.922,5.193 21.79,5.099 Z" />
                            <polygon points="8 21.764 14 17.479 14 0.25 8 4.535" />
                        </g>
                    </g>
                </g>
            );
        case "grid":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g fill="currentColor">
                        <rect x="0" y="0" width="6" height="6" />
                        <rect x="9" y="0" width="6" height="6" />
                        <rect x="18" y="0" width="6" height="6" />
                        <rect x="18" y="9" width="6" height="6" />
                        <rect x="9" y="9" width="6" height="6" />
                        <rect x="0" y="9" width="6" height="6" />
                        <rect x="0" y="18" width="6" height="6" />
                        <rect x="9" y="18" width="6" height="6" />
                        <rect x="18" y="18" width="6" height="6" />
                    </g>
                </g>
            );
        case "email":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g>
                        <g
                            transform="translate(1, 4)"
                            fill="currentColor"
                            fillRule="nonzero"
                        >
                            <path d="M11.002,8.36 L21.097,0.33 C20.782,0.123 20.406,0 20.001,0 L2.001,0 C1.614,0 1.255,0.115 0.948,0.307 L11.002,8.36 Z" />
                            <path d="M21.764,1.076 L11.296,9.391 C11.118,9.525 10.88,9.525 10.702,9.39 L0.26,1.036 C0.1,1.323 0.001,1.648 0,2 L0,13 C0.001,14.103 0.898,15 2.001,15 L19.999,15 C21.102,15 21.999,14.103 22,13 L22,2 C22,1.665 21.91,1.353 21.764,1.076 Z" />
                        </g>
                    </g>
                </g>
            );
        case "phone":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g>
                        <path
                            d="M22.7450921,17.4415 L20.2920921,14.9885 C19.2910921,13.9865 17.6700921,13.9865 16.6670921,14.9885 L16.0100921,15.6455 C13.4400921,13.5195 10.9720921,11.0515 8.85109211,8.4855 L9.50809211,7.8285 C10.5070921,6.8295 10.5070921,5.2035 9.50809211,4.2045 L7.05509211,1.7515 C6.05409211,0.7495 4.43209211,0.7495 3.43009211,1.7515 L2.08409211,3.0975 C0.841092105,4.3395 0.643092105,6.3085 1.61409211,7.7775 C5.54809211,13.7245 10.7710921,18.9475 16.7190921,22.8815 C18.1690921,23.8405 20.1280921,23.6845 21.3990921,22.4115 L22.7450921,21.0655 C23.7490921,20.0645 23.7450921,18.4405 22.7450921,17.4415 Z"
                            fill="currentColor"
                            fillRule="nonzero"
                        />
                    </g>
                </g>
            );
        case "bubbleChat":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g>
                        <path
                            d="M23.5,1 L0.5,1 C0.224,1 0,1.224 0,1.5 L0,17.5 C0,17.776 0.224,18 0.5,18 L7,18 L7,21.5 C7,21.944 7.54,22.168 7.853,21.854 L11.707,18 L23.5,18 C23.776,18 24,17.776 24,17.5 L24,1.5 C24,1.224 23.776,1 23.5,1 Z"
                            fill="currentColor"
                            fillRule="nonzero"
                        />
                    </g>
                </g>
            );
        case "triangleDown":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g>
                        <path
                            d="M23.9348242,2.255 C23.8468242,2.098 23.6808242,2 23.4998242,2 L0.499824176,2 C0.318824176,2 0.152824176,2.098 0.0638241758,2.255 C-0.0241758242,2.412 -0.0211758242,2.605 0.0728241758,2.759 L11.5728242,21.759 C11.6628242,21.908 11.8258242,22 11.9998242,22 C12.1748242,22 12.3368242,21.908 12.4268242,21.759 L23.9268242,2.759 C24.0208242,2.605 24.0248242,2.412 23.9348242,2.255"
                            fill="currentColor"
                        />
                    </g>
                </g>
            );
        case "search":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g fill="currentColor" fillRule="nonzero">
                        <path d="M23.8125,20.1163636 L18.5125,14.7615584 C19.375,13.2592208 19.825,11.5885714 19.825,9.90545455 C19.825,4.44467532 15.375,0 9.9125,0 C4.45,3.54348585e-15 0,4.44467532 0,9.90545455 C0,15.3662338 4.45,19.8109091 9.9125,19.8109091 C11.65625,19.8109091 13.38125,19.3309091 14.91875,18.4145455 L20.19375,23.7506494 C20.3125,23.8690909 20.48125,23.9438961 20.65,23.9438961 C20.81875,23.9438961 20.9875,23.8753247 21.10625,23.7506494 L23.8125,21.0202597 C24.0625,20.7646753 24.0625,20.3657143 23.8125,20.1163636 Z M9.9125,3.8587013 C13.25,3.8587013 15.9625,6.57038961 15.9625,9.90545455 C15.9625,13.2405195 13.25,15.9522078 9.9125,15.9522078 C6.575,15.9522078 3.8625,13.2405195 3.8625,9.90545455 C3.8625,6.57038961 6.575,3.8587013 9.9125,3.8587013 Z" />
                    </g>
                </g>
            );
        case "chevronRight":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <path
                        d="M19.8322393,11.6200895 L6.94763543,0.122172623 C6.84969421,0.035173252 6.72169182,-0.00782643709 6.59465915,0.00117349784 C6.46568705,0.0121734183 6.3464121,0.0731729773 6.26204689,0.17517224 L4.11606749,2.75715357 C4.03267199,2.85715285 3.99097424,2.9891519 4.00164111,3.12115094 C4.01230798,3.25414998 4.07243031,3.37714909 4.17134124,3.46314847 L13.9169775,11.9980868 L4.17134124,20.537025 C4.07243031,20.6230244 4.01230798,20.7460235 4.00164111,20.8790226 C3.99097424,21.0110216 4.03267199,21.1430207 4.11606749,21.2430199 L6.26204689,23.8250013 C6.3464121,23.9270005 6.46568705,23.9880001 6.59465915,23.999 C6.60629573,24 6.61793231,24 6.63053861,24 C6.7459347,24 6.85842165,23.9570003 6.94763543,23.8780009 L19.8322393,12.376084 C19.9379382,12.2810847 20,12.1440857 20,11.9980868 C20,11.8520878 19.9379382,11.7150888 19.8322393,11.6200895"
                        fill="currentColor"
                    />
                </g>
            );
        case "closeCircle":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <path
                        d="M11.5,0 C5.148,0 0,5.148 0,11.5 C0,17.852 5.148,23 11.5,23 C17.852,23 23,17.852 23,11.5 C23,5.148 17.852,0 11.5,0 Z M16.096,15.389 C16.291,15.584 16.291,15.901 16.096,16.096 C15.901,16.291 15.584,16.291 15.389,16.096 L11.5,12.207 L7.611,16.096 C7.416,16.291 7.099,16.291 6.904,16.096 C6.709,15.901 6.709,15.584 6.904,15.389 L10.793,11.5 L6.904,7.611 C6.709,7.416 6.709,7.099 6.904,6.904 C7.099,6.709 7.416,6.709 7.611,6.904 L11.5,10.793 L15.389,6.904 C15.584,6.709 15.901,6.709 16.096,6.904 C16.291,7.099 16.291,7.416 16.096,7.611 L12.207,11.5 L16.096,15.389 Z"
                        fill="currentColor"
                        fillRule="nonzero"
                    />
                </g>
            );
        case "wifiCheck":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g stroke="currentColor" strokeLinejoin="round">
                        <g>
                            <g transform="translate(13, 13)">
                                <polyline
                                    strokeLinecap="round"
                                    points="8 4 5 7 3 5"
                                />
                                <circle cx="5.5" cy="5.5" r="5" />
                            </g>
                            <g strokeLinecap="round">
                                <path d="M3.175,7.629 C6.666,4.137 12.33,4.131 15.827,7.629" />
                                <path d="M18.358,4.203 C13.468,-0.686 5.54,-0.693 0.642,4.203" />
                                <path d="M5.704,11.054 C7.8,8.959 11.198,8.955 13.297,11.054" />
                                <path d="M15.827,7.629 C12.336,4.137 6.671,4.131 3.175,7.629" />
                                <path d="M13.297,11.054 C11.2,8.959 7.804,8.955 5.704,11.054" />
                                <circle cx="9.5" cy="15.744" r="1.79" />
                            </g>
                        </g>
                    </g>
                </g>
            );
        case "checkShield":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g transform="translate(-2, 0)">
                        <g
                            transform="translate(2, 0)"
                            stroke="currentColor"
                            strokeLinejoin="round"
                        >
                            <g>
                                <g>
                                    <polyline
                                        strokeLinecap="round"
                                        points="15.501 7 8.001 14 5.501 11.5"
                                    />
                                    <path d="M0.5,0.5 L19.5,0.5 L19.5,7.764 C19.5,14.367 15.841,20.425 10,23.5 C4.156,20.425 0.5,14.367 0.5,7.764 L0.5,0.5 Z" />
                                </g>
                            </g>
                        </g>
                    </g>
                </g>
            );
        case "devices":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g>
                        <g
                            transform="translate(3, 5)"
                            stroke="currentColor"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                        >
                            <path d="M5.5,0.5 L16,0.5 C16.829,0.5 17.5,1.211 17.5,2.088 L17.5,7.5" />
                            <path d="M12.5,12.5 L2,12.5 C1.172,12.5 0.5,11.789 0.5,10.912 L0.5,6.5" />
                            <path d="M5.5,14.5 L12.5,14.5" />
                            <path d="M9,12.5 L9,14.5" />
                            <path d="M0.5,9.5 L12.5,9.5" />
                        </g>
                        <path
                            d="M6.5,8.5 C6.5,9.05 6.05,9.5 5.5,9.5 L1.5,9.5 C0.95,9.5 0.5,9.05 0.5,8.5 L0.5,1.5 C0.5,0.95 0.95,0.5 1.5,0.5 L5.5,0.5 C6.05,0.5 6.5,0.95 6.5,1.5 L6.5,8.5 Z"
                            stroke="currentColor"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                        />
                        <path
                            d="M23.5,22.5 C23.5,23.05 23.05,23.5 22.5,23.5 L18.5,23.5 C17.95,23.5 17.5,23.05 17.5,22.5 L17.5,15.5 C17.5,14.95 17.95,14.5 18.5,14.5 L22.5,14.5 C23.05,14.5 23.5,14.95 23.5,15.5 L23.5,22.5 Z"
                            stroke="currentColor"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                        />
                        <path
                            d="M17.5,21.5 L23.5,21.5"
                            stroke="currentColor"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                        />
                        <polygon
                            fill="currentColor"
                            fillRule="nonzero"
                            points="4 7 3 7 3 8 4 8"
                        />
                        <path
                            d="M17.5,16.5 L23.5,16.5"
                            stroke="currentColor"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                        />
                    </g>
                </g>
            );
        case "earth":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g fill="currentColor" fillRule="nonzero">
                        <path d="M21.14,4.236 C20.67,5.568 19.714,7.701 18.224,8.447 C18.117,8.5 17.994,8.513 17.879,8.484 C16.822,8.221 15.752,8.611 15.205,8.923 C15.435,9.232 15.692,9.769 15.938,10.697 C16.151,10.772 16.566,10.659 16.775,10.553 C16.968,10.456 17.2,10.494 17.353,10.647 C18.566,11.86 16.994,13.494 16.054,14.472 C15.888,14.644 15.676,14.864 15.558,15.008 C15.598,15.045 15.645,15.087 15.682,15.118 C15.91,15.317 16.223,15.589 16.239,16.02 C16.249,16.307 16.119,16.587 15.853,16.854 C15.334,17.372 14.811,17.677 14.486,17.835 C14.343,19.535 13.09,20.5 11,20.5 C10.024,20.5 9,18.007 9,17.5 C9,17.132 9.16,16.81 9.303,16.526 C9.399,16.333 9.5,16.132 9.5,16 C9.477,15.829 9.07,15.279 8.646,14.854 C8.553,14.761 8.5,14.633 8.5,14.5 C8.5,14.071 8.423,13.78 8.271,13.637 C8.019,13.4 7.422,13.434 6.793,13.47 C6.534,13.484 6.267,13.5 6,13.5 C4.418,13.5 4,11.865 4,11 C4,10.84 4.033,7.083 6.901,6.51 C8.209,6.248 9.091,6.325 9.604,6.746 C9.785,6.894 9.894,7.068 9.95,7.23 C10.476,7.619 11.42,7.398 12.259,7.204 C12.536,7.139 12.805,7.077 13.054,7.039 C13.147,6.194 13.15,5.376 13.063,5.021 C12.45,5.278 11.841,5.276 11.365,5.012 C10.872,4.738 10.565,4.217 10.503,3.545 C10.369,2.107 13.144,0.989 15.05,0.393 C14.074,0.137 13.053,0 12,0 C5.383,0 0,5.383 0,12 C0,18.617 5.383,24 12,24 C18.616,24 24,18.617 24,12 C24,9.042 22.921,6.332 21.14,4.236 Z" />
                    </g>
                </g>
            );
        case "hourglass":
            return (
                <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g fill="currentColor" fillRule="nonzero">
                        <path
                            d="M17.5,23 L16,23 L16,18.5 C16,15.561 14.178,13.045 11.607,12.009 C13.915,11.077 16,8.657 16,5.5 L16,1 L17.5,1 C17.776,1 18,0.776 18,0.5 C18,0.224 17.776,0 17.5,0 L0.5,0 C0.224,0 0,0.224 0,0.5 C0,0.776 0.224,1 0.5,1 L2,1 L2,5.5 C2,8.673 4.106,11.083 6.393,12.009 C3.822,13.045 2,15.561 2,18.5 L2,23 L0.5,23 C0.224,23 0,23.224 0,23.5 C0,23.776 0.224,24 0.5,24 L17.5,24 C17.776,24 18,23.776 18,23.5 C18,23.224 17.776,23 17.5,23 Z M4.137,9 C3.424,8.014 3,6.807 3,5.5 L3,1 L15,1 L15,5.5 C15,6.807 14.576,8.014 13.864,9 L4.137,9 Z M15,19 L12.207,19 L9.354,16.146 C9.159,15.951 8.842,15.951 8.647,16.146 L5.793,19 L3,19 L3,18.5 C3,15.191 5.692,12.5 9,12.5 C12.308,12.5 15,15.191 15,18.5 L15,19 Z"
                            id="Shape"
                        />
                    </g>
                </g>
            );
        default:
            return null;
    }
    /* eslint-enable default-case */
};

export default Glyph;
