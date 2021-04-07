const cacheName = "app-" + "10299e9db643bd34ae7e6b6cc8f175b3ca72ca14";

self.addEventListener("install", event => {
  console.log("installing app worker 10299e9db643bd34ae7e6b6cc8f175b3ca72ca14");
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
  console.log("app worker 10299e9db643bd34ae7e6b6cc8f175b3ca72ca14 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
