import React, { Component } from "react";
import { connect } from "react-redux";
import Downshift from "downshift";
import { Flex } from "grid-styled";

import Absolute from "../Absolute";
import Relative from "../Relative";
import Icon from "../Icon";
import Box from "../Box";
import Image from "../Image";
import Clickable from "../Clickable";
import BoxHiddenOverflow from "../BoxHiddenOverflow";
import FlexFillHeight from "../FlexFillHeight";
import Span from "../Span";

import { fetchResources } from "../../actions/resources/actions";

function onChange(selectedItem) {}

class ResourcesDropdown extends Component {
    componentDidMount() {
        this.props.fetchResources();
    }

    renderData() {
        if (this.props.resources.error) {
            return (
                <div>
                    An error has occurred:{" "}
                    {this.props.resources.error.toString()}
                </div>
            );
        } else if (!this.props.resources.pending) {
            return (
                <BoxHiddenOverflow
                    bg="white"
                    borderRadius={4}
                    w={1}
                    boxShadow="0px 8px 20px rgba(0, 0, 0, 0.1)"
                >
                    <Flex flexWrap="wrap" p={4} alignItems="center">
                        {this.props.resources.data.map(item => {
                            return (
                                <Box w={1 / 3} key={item.name} px={4}>
                                    <a
                                        href={item.link}
                                        rel="noopener noreferrer"
                                        target="_blank"
                                    >
                                        <Relative w={1} pt="100%">
                                            <Absolute
                                                top={0}
                                                left={0}
                                                bottom={0}
                                                right={0}
                                            >
                                                <FlexFillHeight alignItems="center">
                                                    <div>
                                                        <Image
                                                            src={item.imageLink}
                                                        />
                                                    </div>
                                                </FlexFillHeight>
                                            </Absolute>
                                        </Relative>
                                        <Box mt={2}>
                                            <Span>{item.name}</Span>
                                        </Box>
                                    </a>
                                </Box>
                            );
                        })}
                    </Flex>
                </BoxHiddenOverflow>
            );
        } else {
            return null;
        }
    }

    render() {
        if (!this.props.resources.pending) {
            return (
                <Downshift
                    onChange={onChange}
                    render={({
                        getLabelProps,
                        getInputProps,
                        getButtonProps,
                        getItemProps,
                        isOpen,
                        toggleMenu,
                        clearSelection,
                        selectedItem,
                        inputValue,
                        highlightedIndex
                    }) => {
                        return (
                            <div>
                                <Clickable onClick={toggleMenu}>
                                    <Icon
                                        glyph="grid"
                                        size={16}
                                        color="white"
                                    />
                                </Clickable>
                                {isOpen ? (
                                    <Relative>
                                        <Absolute
                                            right={0}
                                            zIndex={1}
                                            top={8}
                                            w={320}
                                        >
                                            {this.renderData()}
                                        </Absolute>
                                    </Relative>
                                ) : null}
                            </div>
                        );
                    }}
                />
            );
        }
        return <div />;
    }
}

function mapStateToProps(state, ownProps) {
    return {
        resources: state.resources
    };
}

const mapDispatchToProps = {
    fetchResources
};

export default connect(mapStateToProps, mapDispatchToProps)(ResourcesDropdown);
