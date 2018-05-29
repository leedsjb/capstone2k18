import React, { Component } from "react";
import { Flex } from "grid-styled";

import Box from "../Box";
import Clickable from "../Clickable";
import DisplayWhenOpen from "../DisplayWhenOpen";
import Icon from "../Icon";
import Span from "../Span";

class AccordionSection extends Component {
    constructor(props) {
        super(props);
        this.state = {
            active: null,
            isopen: false
        };
    }

    toggleisopen = () => {
        this.setState({
            active: this.state.active ? null : this.props.title,
            isopen: !this.state.isopen
        });
    };

    render() {
        return (
            <div>
                <Clickable onClick={this.toggleisopen}>
                    <Box borderBottom="1px solid black" p={3}>
                        <Flex justifyContent="space-between">
                            <Span fontWeight="bold"> {this.props.title}</Span>
                            <Icon
                                glyph={
                                    this.state.active === this.props.title
                                        ? "chevronUp"
                                        : "chevronDown"
                                }
                                size={16}
                                color="black"
                            />
                        </Flex>
                    </Box>
                </Clickable>
                <DisplayWhenOpen
                    isopen={this.state.isopen ? 1 : 0}
                    borderBottom={
                        this.state.isopen ? "1px solid black" : "none"
                    }
                    px={3}
                >
                    {this.props.children}
                </DisplayWhenOpen>
            </div>
        );
    }
}

export default AccordionSection;
