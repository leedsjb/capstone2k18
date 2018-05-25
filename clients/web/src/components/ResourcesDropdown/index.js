import React from "react";
import Downshift from "downshift";

import Absolute from "../Absolute";
import Relative from "../Relative";
import Icon from "../Icon";
import Image from "../Image";

import ResourcesProvider from "../../containers/ResourcesProvider";

function onChange(selectedItem) {}

const ResourcesDropdown = () => {
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
                                                {resources.pending ? null : (
                                                    <div>
                                                        {resources.data[0].name}
                                                        <Image
                                                            src={
                                                                resources
                                                                    .data[0]
                                                                    .imageLink
                                                            }
                                                        />
                                                    </div>
                                                )}
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

export default ResourcesDropdown;
