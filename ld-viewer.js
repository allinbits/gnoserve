class LDViewer extends HTMLElement {
    constructor() {
        super();
        this.attachShadow({mode: 'open'});
        this.shadowRoot.innerHTML = `
<style>
:host { --bg:#0b1020; --fg:#e6eefc; --muted:#a9b4d0; --card:#121a33; --accent:#8abdff; display:block }
*{box-sizing:border-box}
body{margin:0}
header{padding:24px; border-bottom:1px solid #1e2a52; background:linear-gradient(180deg,#0e1530,transparent)}
h1{font-size:1.4rem; margin:0 0 6px}
p.lead{margin:0; color:var(--muted)}
main{display:grid; gap:16px; grid-template-columns: 1.1fr 1fr; padding:16px}
.col{display:flex; flex-direction:column; gap:12px}
.card{background:var(--card); border:1px solid #1e2a52; border-radius:14px; padding:14px}
label b{display:block; margin-bottom:6px; color:#cfe1ff}
input[type="text"], textarea{width:100%; background:#0b1330; color:#e6eefc; border:1px solid #2a3a7a; border-radius:10px; padding:10px 12px; font-family:ui-monospace,SFMono-Regular,Menlo,Consolas,monospace}
textarea{min-height:130px}
.row{display:flex; gap:10px; flex-wrap:wrap}
button{border:1px solid #2a3a7a; background:#0f1a45; color:#e8f0ff; padding:10px 14px; border-radius:12px; cursor:pointer}
button:hover{border-color:#3a5ad2}
.btn-accent{background:#1a2e75; border-color:#3157ff}
.btn-ghost{background:transparent}
.small{font-size:12px; color:var(--muted)}
pre{margin:0; white-space:pre-wrap; word-break:break-word; max-height:360px; overflow:auto; font-size:13px; background:#0b1330; border:1px solid #2a3a7a; border-radius:10px; padding:12px}
.status{display:grid; grid-template-columns:1fr 1fr; gap:10px}
.kv{display:grid; grid-template-columns:max-content 1fr; gap:6px 10px; align-items:center}
.kv div:first-child{color:#b7c8ff}
.mono{font-family:ui-monospace,SFMono-Regular,Menlo,Consolas,monospace}
.error{color:#ffb3c7}
footer{padding:16px; color:var(--muted)}
@media (max-width: 1000px){ main{grid-template-columns:1fr} }
</style>
<header>
    <h1>Portable JSON-LD File Builder</h1>
    <p class="lead">Type JSON, add an optional <span class="mono">@context</span>, preview expanded/flattened/canonical forms, download files, and compute a content hash & CID (raw) of the canonical N-Quads.</p>
</header>
<main>
    <section class="col">
        <div class="card">
            <label>
                <b>JSON-LD Context (optional)</b>
                <input id="context" type="text" value='{"@vocab":"https://schema.org/"}' spellcheck="false"/>
                <div class="small">Tip: Enter JSON (e.g., <span class="mono">{ "@vocab":"https://schema.org/" }</span>) or a URL (left as-is, not fetched).</div>
            </label>
        </div>
        <div class="card">
            <label>
                <b>Data (JSON object)</b>
                <textarea id="data" spellcheck="false">{ "@id": "urn:uuid:1234", "@type": "Person", "name": "Alice" }</textarea>
            </label>
            <div class="row">
                <button class="btn-accent" id="btn-download-jsonld">Download JSON-LD</button>
                <button id="btn-download-nq">Download N-Quads</button>
                <button id="btn-compact">Compact</button>
                <button id="btn-expand">Expand</button>
                <button id="btn-flatten">Flatten</button>
                <button id="btn-frame">Apply Frame</button>
                <button id="btn-canon">Canonicalize</button>
            </div>
            <div class="row">
                <label class="small"><input type="checkbox" id="opt-pretty" checked/> Pretty-print</label>
                <label class="small"><input type="checkbox" id="opt-sort"/> Sort keys (stable)</label>
                <label class="small"><input type="checkbox" id="opt-keepctx" checked/> Keep <span class="mono">@context</span> in output</label>
            </div>
        </div>
        <div class="card">
            <label>
                <b>Frame (optional, JSON object)</b>
                <textarea id="frame" placeholder='{"@type":"Person"}' spellcheck="false"></textarea>
            </label>
            <div class="small">Frames reshape data for specific views. Only used when you click <i>Apply Frame</i>.</div>
        </div>
    </section>
    <section class="col">
        <div class="card">
            <b>Output</b>
            <pre id="output" aria-live="polite"></pre>
            <div class="small" id="msg"></div>
        </div>
        <div class="card">
            <b>Canonicalization & Integrity</b>
            <div class="status">
                <div class="kv">
                    <div>N-Quads bytes:</div><div class="mono" id="stat-bytes">–</div>
                    <div>SHA-256 (hex):</div><div class="mono" id="stat-sha">–</div>
                    <div>Multihash (base32):</div><div class="mono" id="stat-mh">–</div>
                    <div>Raw CIDv1 (base32):</div><div class="mono" id="stat-cid">–</div>
                </div>
                <div>
                    <button id="btn-copy-hash" class="btn-ghost">Copy hash</button>
                    <button id="btn-copy-cid" class="btn-ghost">Copy CID</button>
                </div>
            </div>
            <div class="small">CID note: This creates a <span class="mono">raw</span> CIDv1 (multicodec <span class="mono">0x55</span>) over the canonical N-Quads bytes. If you need <span class="mono">dag-json</span> CIDs, hash canonical DAG-JSON bytes instead.</div>
        </div>
    </section>
</main>
<footer>
    <span class="small">All processing happens locally in your browser. No network requests are made except library CDNs.</span>
</footer>
        `;
    }

    connectedCallback() {
        const $ = sel => this.shadowRoot.querySelector(sel);
        const out = $('#output');
        const msg = $('#msg');
        const pretty = () => $('#opt-pretty').checked;
        const sortKeys = () => $('#opt-sort').checked;
        const keepCtx = () => $('#opt-keepctx').checked;

        // Load initial data from <script type="application/ld+json"> if present
        let initialData = '{ "@id": "urn:uuid:1234", "@type": "Person", "name": "Alice" }';
        const script = this.querySelector('script[type="application/ld+json"]');
        if (script) {
            try {
                const json = JSON.parse(script.textContent);
                initialData = JSON.stringify(json, null, 2);
            } catch (e) {
                // fallback to default if parsing fails
            }
        }
        this.shadowRoot.querySelector('#data').value = initialData;

        function parseContext(raw){
            if(!raw) return undefined;
            const trimmed = raw.trim();
            if(!trimmed) return undefined;
            if((trimmed.startsWith('{') && trimmed.endsWith('}')) || (trimmed.startsWith('[') && trimmed.endsWith(']'))){
                return JSON.parse(trimmed);
            }
            return trimmed;
        }

        function stableStringify(value){
            if(!sortKeys()) return JSON.stringify(value, null, pretty()?2:0);
            const seen = new WeakSet();
            function sort(obj){
                if(obj && typeof obj === 'object'){
                    if(seen.has(obj)) return obj;
                    seen.add(obj);
                    if(Array.isArray(obj)) return obj.map(sort);
                    return Object.fromEntries(Object.keys(obj).sort().map(k => [k, sort(obj[k])]));
                }
                return obj;
            }
            return JSON.stringify(sort(value), null, pretty()?2:0);
        }

        function assembleInput(){
            const ctxRaw = $('#context').value;
            const dataRaw = $('#data').value;
            let data;
            try { data = JSON.parse(dataRaw); }
            catch(err){ throw new Error('Data is not valid JSON: '+ err.message); }
            const ctx = parseContext(ctxRaw);
            if(ctx !== undefined){ data['@context'] = ctx; }
            else { delete data['@context']; }
            return data;
        }

        async function compactForOutput(doc){
            if(!keepCtx()) return doc;
            const ctx = doc['@context'];
            if(!ctx) return doc;
            try{ return await window.jsonld.compact(doc, ctx); }
            catch{ return doc; }
        }

        function download(text, filename, type){
            const blob = new Blob([text], {type});
            window.saveAs(blob, filename);
        }

        async function toCanonicalNQ(doc){
            return await window.jsonld.canonize(doc, { algorithm: 'URDNA2015', format: 'application/n-quads' });
        }

        async function sha256(bytes){
            const digest = await crypto.subtle.digest('SHA-256', bytes);
            return new Uint8Array(digest);
        }

        function toHex(bytes){
            return [...bytes].map(b=>b.toString(16).padStart(2,'0')).join('');
        }

        function textEncoder(){ return new TextEncoder(); }

        async function updateIntegrity(nquads){
            try{
                const bytes = textEncoder().encode(nquads);
                $('#stat-bytes').textContent = String(bytes.length);
                const digest = await sha256(bytes);
                $('#stat-sha').textContent = toHex(digest);

                const mf = window.multiformats;
                const mh = await mf.hasher.sha256.digest(bytes);
                const mhB32 = mf.bases.base32.encode(mh.bytes);
                $('#stat-mh').textContent = mhB32;

                const raw = mf.multicodec.get('raw');
                const cid = mf.CID.createV1(raw.code, mh);
                $('#stat-cid').textContent = cid.toString(mf.bases.base32);

                return { mhB32, cid: cid.toString(mf.bases.base32) };
            }catch(err){
                $('#stat-bytes').textContent = '–';
                $('#stat-sha').textContent = '–';
                $('#stat-mh').textContent = '–';
                $('#stat-cid').textContent = '–';
                console.error(err);
                return null;
            }
        }

        function show(value){
            out.textContent = typeof value === 'string' ? value : stableStringify(value);
        }

        function note(text, isError=false){
            msg.textContent = text;
            msg.className = isError ? 'small error' : 'small';
        }

        $('#btn-compact').addEventListener('click', async ()=>{
            try{
                note('');
                const doc = assembleInput();
                const ctx = doc['@context'] ?? {"@vocab":"https://schema.org/"};
                const compacted = await window.jsonld.compact(doc, ctx);
                show(compacted);
            }catch(err){ note(err.message, true); }
        });

        $('#btn-expand').addEventListener('click', async ()=>{
            try{
                note('');
                const doc = assembleInput();
                const expanded = await window.jsonld.expand(doc);
                show(expanded);
            }catch(err){ note(err.message, true); }
        });

        $('#btn-flatten').addEventListener('click', async ()=>{
            try{
                note('');
                const doc = assembleInput();
                const flattened = await window.jsonld.flatten(doc);
                show(flattened);
            }catch(err){ note(err.message, true); }
        });

        $('#btn-frame').addEventListener('click', async ()=>{
            try{
                note('');
                const doc = assembleInput();
                let frame = $('#frame').value.trim();
                if(!frame) { note('No frame provided.'); return; }
                frame = JSON.parse(frame);
                const framed = await window.jsonld.frame(doc, frame);
                show(framed);
            }catch(err){ note(err.message, true); }
        });

        $('#btn-canon').addEventListener('click', async ()=>{
            try{
                note('');
                const doc = assembleInput();
                const nquads = await toCanonicalNQ(doc);
                show(nquads);
                await updateIntegrity(nquads);
            }catch(err){ note(err.message, true); }
        });

        $('#btn-download-jsonld').addEventListener('click', async ()=>{
            try{
                note('');
                const doc = assembleInput();
                const outDoc = await compactForOutput(doc);
                const text = stableStringify(outDoc);
                download(text, 'document.jsonld', 'application/ld+json');
            }catch(err){ note(err.message, true); }
        });

        $('#btn-download-nq').addEventListener('click', async ()=>{
            try{
                note('');
                const doc = assembleInput();
                const nquads = await toCanonicalNQ(doc);
                download(nquads, 'document.nq', 'application/n-quads');
                await updateIntegrity(nquads);
            }catch(err){ note(err.message, true); }
        });

        $('#btn-copy-hash').addEventListener('click', async ()=>{
            const val = $('#stat-sha').textContent.trim();
            if(!val || val==='–') return;
            await navigator.clipboard.writeText(val);
            note('SHA-256 copied to clipboard.');
        });

        $('#btn-copy-cid').addEventListener('click', async ()=>{
            const val = $('#stat-cid').textContent.trim();
            if(!val || val==='–') return;
            await navigator.clipboard.writeText(val);
            note('CID copied to clipboard.');
        });

        (async ()=>{
            try{
                const doc = assembleInput();
                const nquads = await toCanonicalNQ(doc);
                show(nquads);
                await updateIntegrity(nquads);
            }catch(err){ note('Ready. Enter JSON and click a button.'); }
        })();
    }
}
customElements.define('ld-viewer', LDViewer);