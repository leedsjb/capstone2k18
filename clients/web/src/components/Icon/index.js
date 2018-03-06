//Based on: https://gist.github.com/uberbryn/ebb3eed5cc29767ed63b8eb73ae86ae5

import React, { Component } from "react";

import InlineSVG from "./InlineSVG";
import SVGWrapper from "./SVGWrapper";
import Glyph from "./Glyph";

class Icon extends Component {
    render() {
        const { size, onClick, glyph } = this.props;

        return (
            <SVGWrapper size={size} className={"icon"} onClick={onClick}>
                <InlineSVG
                    fillRule="evenodd"
                    clipRule="evenodd"
                    strokeLinejoin="round"
                    strokeMiterlimit="1.414"
                    xmlns="http://www.w3.org/2000/svg"
                    aria-labelledby="title"
                    viewBox="0 0 32 32"
                    preserveAspectRatio="xMidYMid meet"
                    fit="true"
                    id={glyph}
                >
                    <title id="title">{glyph}</title>
                    <Glyph glyph={glyph} />
                </InlineSVG>
            </SVGWrapper>
        );
    }
}

Icon.defaultProps = {
    size: 32
};

export default Icon;
