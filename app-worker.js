const cacheName = "app-" + "d8570929835ee35ce30ee5410e1cafd05214c079";

self.addEventListener("install", event => {
  console.log("installing app worker d8570929835ee35ce30ee5410e1cafd05214c079");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
      return cache.addAll([
        "",
        "/liwasc",
        "/liwasc/app.css",
        "/liwasc/app.js",
        "/liwasc/manifest.webmanifest",
        "/liwasc/wasm_exec.js",
        "/liwasc/web/app.wasm",
        "/liwasc/web/icon.png",
        "/liwasc/web/index.css",
        "https://unpkg.com/@patternfly/patternfly@4.90.5/patternfly-addons.css",
        "https://unpkg.com/@patternfly/patternfly@4.90.5/patternfly.css",
        
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
  console.log("app worker d8570929835ee35ce30ee5410e1cafd05214c079 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
