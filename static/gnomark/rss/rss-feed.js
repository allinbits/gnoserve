class RssFeed extends HTMLElement {
    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
    }

    connectedCallback() {
        const data = this.getFeedData();
        if (data) {
            this.render(data);
        } else {
            this.shadowRoot.innerHTML = `<p>Error: No feed data found.</p>`;
        }
    }

    getFeedData() {
        try {
            const rawData = this.textContent.trim();
            return JSON.parse(rawData).feed;
        } catch (e) {
            console.error('Invalid JSON data in <rss-feed>', e);
            return null;
        }
    }

    render(feed) {
        const { title, link, description, items } = feed;

        const template = `
            <style>
                .rss-feed { font-family: Arial, sans-serif; line-height: 1.5; }
                .rss-title { font-size: 1.5em; font-weight: bold; }
                .rss-description { margin-bottom: 1em; color: #555; }
                .rss-item { margin-bottom: 1em; }
                .rss-item a { text-decoration: none; color: #007BFF; }
                .rss-item a:hover { text-decoration: underline; }
            </style>
            <div class="rss-feed">
                <div class="rss-title"><a href="${link}" target="_blank">${title}</a></div>
                <div class="rss-description">${description}</div>
                <div class="rss-items">
                    ${items.map(item => `
                        <div class="rss-item">
                            <a href="${item.link}" target="_blank">${item.title}</a>
                            <p>${item.description}</p>
                        </div>
                    `).join('')}
                </div>
            </div>
        `;

        this.shadowRoot.innerHTML = template;
    }
}

if (!customElements.get('rss-feed')) {
    customElements.define('rss-feed', RssFeed);
}