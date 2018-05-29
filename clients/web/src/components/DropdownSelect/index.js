import React, { Component } from "react";
import { withTheme } from "styled-components";
import Downshift from "downshift";

import Absolute from "../Absolute";
import Relative from "../Relative";
import ButtonDropdown from "../ButtonDropdown";
import Span from "../Span";
import DropdownItem from "../DropdownItem";
import BoxHiddenOverflow from "../BoxHiddenOverflow";
import Clickable from "../Clickable";

class DropdownSelect extends Component {
    render() {
        let { items, onChange } = this.props;

        return (
            <Downshift
                onChange={onChange}
                defaultSelectedItem={items[0]}
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
                }) => (
                    <div>
                        <ButtonDropdown onClick={toggleMenu}>
                            {selectedItem}
                        </ButtonDropdown>
                        {isOpen ? (
                            <Relative zIndex={1}>
                                <Absolute
                                    top={8}
                                    left={0}
                                    right={0}
                                    minWidth={160}
                                >
                                    <BoxHiddenOverflow
                                        bg="white"
                                        borderRadius={4}
                                        w={1}
                                        boxShadow={
                                            this.props.theme.boxShadows.mid
                                        }
                                        border={`1px solid ${
                                            this.props.theme.colors.gray5
                                        }`}
                                    >
                                        {items.map((item, i) => {
                                            return (
                                                <div>
                                                    <Clickable>
                                                        <DropdownItem
                                                            {...getItemProps({
                                                                item
                                                            })}
                                                            key={item}
                                                        >
                                                            <Span
                                                                fontWeight={
                                                                    selectedItem ===
                                                                    item
                                                                        ? "bold"
                                                                        : "normal"
                                                                }
                                                            >
                                                                {item}
                                                            </Span>
                                                        </DropdownItem>
                                                    </Clickable>
                                                </div>
                                            );
                                        })}
                                    </BoxHiddenOverflow>
                                </Absolute>
                            </Relative>
                        ) : null}
                    </div>
                )}
            />
        );
    }
}

export default withTheme(DropdownSelect);
