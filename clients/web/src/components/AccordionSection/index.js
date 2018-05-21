import React, { Component } from "react";

import Clickable from "../Clickable";
import DisplayWhenOpen from "../DisplayWhenOpen";
import Box from "../Box";
import Span from "../Span";

class AccordionSection extends Component {
    constructor(props) {
        super(props);
        this.state = {
            isopen: false
        };
    }

    toggleisopen = () => {
        this.setState({
            isopen: !this.state.isopen
        });
    };

    render() {
        return (
            <div>
                <Clickable onClick={this.toggleisopen}>
                    <Box
                        borderTop={
                            this.props.id === 0 ? "1px solid black" : "none"
                        }
                        borderBottom="1px solid black"
                    >
                        <Span fontWeight="bold"> {this.props.title}</Span>
                    </Box>
                </Clickable>
                <DisplayWhenOpen
                    isopen={this.state.isopen}
                    borderBottom={
                        this.state.isopen ? "1px solid black" : "none"
                    }
                >
                    {this.props.children}
                </DisplayWhenOpen>
            </div>
        );
    }
}

export default AccordionSection;
