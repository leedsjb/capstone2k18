import React from "react";
import Downshift from "downshift";

import Absolute from "../Absolute";

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
                    <button onClick={toggleMenu}>{selectedItem}</button>
                    {isOpen ? (
                        <Absolute bg="white">
                            {items.map(item => (
                                <div {...getItemProps({ item })} key={item}>
                                    {item}
                                </div>
                            ))}
                        </Absolute>
                    ) : null}
                </div>
            )}
        />
    );
};

export default DropdownSelect;
