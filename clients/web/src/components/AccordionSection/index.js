import React, { Component } from "react";
import { withTheme } from "styled-components";
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
                    <Box
                        borderBottom={`1px solid ${
                            this.props.theme.colors.gray4
                        }`}
                        px={4}
                        py={3}
                    >
                        <Flex justifyContent="space-between">
                            <Span fontWeight="bold"> {this.props.title}</Span>
                            <Box>
                                <Icon
                                    glyph={
                                        this.state.active === this.props.title
                                            ? "chevronUp"
                                            : "chevronDown"
                                    }
                                    color="black1"
                                    size={16}
                                />
                            </Box>
                        </Flex>
                    </Box>
                </Clickable>
                <DisplayWhenOpen
                    isopen={this.state.isopen ? 1 : 0}
                    borderBottom={`1px solid ${this.props.theme.colors.gray4}`}
                    px={4}
                >
                    <Box my={6}>{this.props.children}</Box>
                </DisplayWhenOpen>
            </div>
        );
    }
}

export default withTheme(AccordionSection);
