import React from "react";
import Downshift from "downshift";

import Absolute from "../Absolute";
import Relative from "../Relative";
import ButtonDropdown from "../ButtonDropdown";
import Span from "../Span";
import DropdownItem from "../DropdownItem";
import BoxHiddenOverflow from "../BoxHiddenOverflow";
import Clickable from "../Clickable";

const DropdownSelect = ({ items, onChange }) => {
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
                            <Absolute top={8} left={0} right={0} minWidth={160}>
                                <BoxHiddenOverflow
                                    bg="white"
                                    borderRadius={4}
                                    w={1}
                                    boxShadow="0px 8px 20px rgba(0, 0, 0, 0.1)"
                                >
                                    {items.map(item => (
                                        <Clickable>
                                            <DropdownItem
                                                {...getItemProps({ item })}
                                                key={item}
                                            >
                                                <Span
                                                    fontWeight={
                                                        selectedItem === item
                                                            ? "bold"
                                                            : "normal"
                                                    }
                                                >
                                                    {item}
                                                </Span>
                                            </DropdownItem>
                                        </Clickable>
                                    ))}
                                </BoxHiddenOverflow>
                            </Absolute>
                        </Relative>
                    ) : null}
                </div>
            )}
        />
    );
};

export default DropdownSelect;
