import React from "react";
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
                                                {resources.pending ? null : (
                                                    <BoxHiddenOverflow
                                                        bg="white"
                                                        borderRadius={4}
                                                        w={1}
                                                        boxShadow="0px 8px 20px rgba(0, 0, 0, 0.1)"
                                                    >
                                                        <Flex
                                                            flexWrap="wrap"
                                                            p={4}
                                                            alignItems="center"
                                                        >
                                                            {resources.data.map(
                                                                item => {
                                                                    return (
                                                                        <Box
                                                                            w={
                                                                                1 /
                                                                                3
                                                                            }
                                                                            key={
                                                                                item.name
                                                                            }
                                                                            px={
                                                                                4
                                                                            }
                                                                        >
                                                                            <a
                                                                                href={
                                                                                    item.link
                                                                                }
                                                                                target="_blank"
                                                                            >
                                                                                <Relative
                                                                                    w={
                                                                                        1
                                                                                    }
                                                                                    pt="100%"
                                                                                >
                                                                                    <Absolute
                                                                                        top={
                                                                                            0
                                                                                        }
                                                                                        left={
                                                                                            0
                                                                                        }
                                                                                        bottom={
                                                                                            0
                                                                                        }
                                                                                        right={
                                                                                            0
                                                                                        }
                                                                                    >
                                                                                        <FlexFillHeight alignItems="center">
                                                                                            <div
                                                                                            >
                                                                                                <Image
                                                                                                    src={
                                                                                                        item.imageLink
                                                                                                    }
                                                                                                />
                                                                                            </div>
                                                                                        </FlexFillHeight>
                                                                                    </Absolute>
                                                                                </Relative>
                                                                                <Box
                                                                                    mt={
                                                                                        2
                                                                                    }
                                                                                >
                                                                                    <Span
                                                                                    >
                                                                                        {
                                                                                            item.name
                                                                                        }
                                                                                    </Span>
                                                                                </Box>
                                                                            </a>
                                                                        </Box>
                                                                    );
                                                                }
                                                            )}
                                                        </Flex>
                                                    </BoxHiddenOverflow>
                                                )}
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

export default ResourcesDropdown;
