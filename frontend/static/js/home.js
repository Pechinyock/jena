function acceptAnswer(num){
    const descusContainer = document.getElementById('descuss-container-' + num);
    const buttons = descusContainer.querySelectorAll('button');

    buttons.forEach(button => {
        button.disabled = true;
    });
}