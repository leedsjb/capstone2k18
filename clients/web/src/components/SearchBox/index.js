import React, { Component } from "react";
import { Flex } from "grid-styled";
import { withTheme } from "styled-components";

import TextInput from "../TextInput";
import Icon from "../Icon";
import Box from "../Box";
import Clickable from "../Clickable";

class SearchBox extends Component {
    handleFocus = () => {
        this.props.handleFocus();
    };

    render() {
        return (
            <Box
                borderRadius={32}
                border={`1px solid ${this.props.theme.colors.border}`}
                bg="white"
                py={1}
            >
                <Flex alignItems="center">
                    <Box px={2}>
                        <Icon glyph="search" size={16} />
                    </Box>
                    <TextInput
                        placeholder="Search"
                        onFocus={this.handleFocus}
                        onChange={event =>
                            this.props.handleChange(event.target.value)
                        }
                        py={0}
                        value={this.props.query}
                    />
                    <div>
                        <Box px={2}>
                            {this.props.isSearching ? (
                                <Clickable
                                    onClick={() => {
                                        this.props.handleClear();
                                    }}
                                >
                                    <Icon glyph="closeCircle" size={16} />
                                </Clickable>
                            ) : null}
                        </Box>
                    </div>
                </Flex>
            </Box>
        );
    }
}

export default withTheme(SearchBox);
