const cacheName = "app-" + "35eceebf9abf0520a988d29541355b161610b7b1";

self.addEventListener("install", event => {
  console.log("installing app worker 35eceebf9abf0520a988d29541355b161610b7b1");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
      return cache.addAll([
        "/liwasc",
        "/liwasc/",
        "/liwasc/app.css",
        "/liwasc/app.js",
        "/liwasc/manifest.webmanifest",
        "/liwasc/wasm_exec.js",
        "/liwasc/web/app.wasm",
        "/liwasc/web/icon.png",
        
      ]);
    })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker 35eceebf9abf0520a988d29541355b161610b7b1 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
