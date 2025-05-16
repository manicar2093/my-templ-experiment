(function(){
	function setLinkConfirmation(id) {
		const link = document.querySelector(id)
		link.addEventListener('click', (event)=>{
			event.preventDefault();
			const method = link.dataset.method;
			const confirmMessage= link.dataset.confirm;

			if(method === 'GET') {
				document.location.href = link.href;
				return;
			}

			if(confirm(confirmMessage)) {
				document.location.href = link.href;
				return
			}

			return;
		})
	}
})()
