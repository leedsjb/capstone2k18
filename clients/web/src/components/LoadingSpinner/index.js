import React from "react";

import "loaders.css/loaders.min.css";

import FlexFillHeight from "../FlexFillHeight";
import ColoredCircle from "./ColoredCircle";

const LoadingSpinner = props => {
    //TODO: Replace this with an actual spinner
    return (
        <FlexFillHeight justifyContent="center" alignItems="center">
            <div className="loader">
                <div className="loader-inner ball-clip-rotate">
                    <ColoredCircle />
                </div>
            </div>
        </FlexFillHeight>
    );
};

export default LoadingSpinner;
