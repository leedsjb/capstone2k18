import React from "react";
import Downshift from "downshift";
import { Link } from "react-router-dom";

import ProfileAvatar from "../ProfileAvatar";
import Absolute from "../Absolute";
import Relative from "../Relative";
import Icon from "../Icon";
import Image from "../Image";

import ResourcesProvider from "../../containers/ResourcesProvider";

function onChange(selectedItem) {
    console.log("Selected item", selectedItem);
}

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
                                    <Relative>
                                        <ProfileAvatar onClick={toggleMenu} />
                                        {isOpen ? (
                                            <Absolute
                                                bg="white"
                                                right={0}
                                                zIndex={1}
                                            >
                                                <Link to="/profile">
                                                    Profile...
                                                </Link>
                                                <div>Sign out</div>
                                            </Absolute>
                                        ) : null}
                                    </Relative>
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
