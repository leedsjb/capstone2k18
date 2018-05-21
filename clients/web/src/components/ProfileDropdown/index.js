import React from "react";
import Downshift from "downshift";
import { Link } from "react-router-dom";

import ProfileAvatar from "../ProfileAvatar";
import Absolute from "../Absolute";
import Relative from "../Relative";
import BoxHiddenOverflow from "../BoxHiddenOverflow";
import DropdownItem from "../DropdownItem";
import Clickable from "../Clickable";

import ResourcesProvider from "../../containers/ResourcesProvider";

function onChange(selectedItem) {}

const ProfileDropdown = () => {
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
                                        <ProfileAvatar fName="Dave" />
                                    </Clickable>
                                    {isOpen ? (
                                        <Relative>
                                            <Absolute
                                                right={0}
                                                zIndex={1}
                                                top={8}
                                                minWidth={120}
                                            >
                                                <BoxHiddenOverflow
                                                    bg="white"
                                                    borderRadius={4}
                                                    w={1}
                                                    boxShadow="0px 8px 20px rgba(0, 0, 0, 0.1)"
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
};

export default ProfileDropdown;
