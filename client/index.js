async function shortenUrl() {
    let url = document.getElementById("urlTextBox").value;

    try {
        const response = await fetch('http://0.0.0.0:8089/create', {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                url:url
            })
        });
        
        if (!response.ok) throw new Error('Request failed');
        const data = await response.json();
        console.log(data);
    } catch (error) {
        console.error('Error:', error);
    }

}
