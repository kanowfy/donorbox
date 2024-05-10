const calculateProgress = (current, goal) => {
    return Math.floor(current / goal * 100);
}

export default { calculateProgress }
