import React, { Component } from "react";
import Downshift from "downshift";
import { Link } from "react-router-dom";
import { withTheme } from "styled-components";

import ColoredAvatar from "../ColoredAvatar";
import Absolute from "../Absolute";
import Relative from "../Relative";
import BoxHiddenOverflow from "../BoxHiddenOverflow";
import DropdownItem from "../DropdownItem";
import Clickable from "../Clickable";

import ResourcesProvider from "../../containers/ResourcesProvider";

function onChange(selectedItem) {}

class ProfileDropdown extends Component {
    render() {
        return (
            <ResourcesProvider
                render={({ resources }) => {
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
                                            <ColoredAvatar fName="Brian" />
                                        </Clickable>
                                        {isOpen ? (
                                            <Relative>
                                                <Absolute
                                                    right={0}
                                                    zIndex={42}
                                                    top={8}
                                                    minWidth={120}
                                                >
                                                    <BoxHiddenOverflow
                                                        bg="white"
                                                        borderRadius={4}
                                                        w={1}
                                                        boxShadow={
                                                            this.props.theme
                                                                .boxShadows.mid
                                                        }
                                                        border={`1px solid ${
                                                            this.props.theme
                                                                .colors.gray5
                                                        }`}
                                                    >
                                                        <Link to="/profile">
                                                            <DropdownItem>
                                                                My profile
                                                            </DropdownItem>
                                                        </Link>
                                                        <DropdownItem>
                                                            Sign out
                                                        </DropdownItem>
                                                    </BoxHiddenOverflow>
                                                </Absolute>
                                            </Relative>
                                        ) : null}
                                    </div>
                                );
                            }}
                        />
                    );
                }}
            />
        );
    }
}

export default withTheme(ProfileDropdown);
