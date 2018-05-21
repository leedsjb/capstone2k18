import React, { Component } from "react";
import { connect } from "react-redux";
import Downshift from "downshift";

import Absolute from "../Absolute";
import Relative from "../Relative";
import Icon from "../Icon";
import Image from "../Image";

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
                <div>
                    {this.props.resources.data[0].name}
                    <Image src={this.props.resources.data[0].imageLink} />
                </div>
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
                                <Relative>
                                    <Icon
                                        glyph="grid"
                                        size={16}
                                        onClick={toggleMenu}
                                        color="white"
                                    />
                                    {isOpen ? (
                                        <Absolute
                                            bg="white"
                                            right={0}
                                            zIndex={1}
                                        >
                                            {this.renderData()}
                                        </Absolute>
                                    ) : null}
                                </Relative>
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
