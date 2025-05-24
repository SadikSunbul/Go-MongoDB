document.addEventListener('DOMContentLoaded', function() {
    // Dil değiştirme işlevselliği
    const langButtons = document.querySelectorAll('.lang-btn');
    let currentLang = 'en'; // Varsayılanı en yap

    function changeLang(lang) {
        currentLang = lang;
        
        // Aktif dil butonunu güncelle
        langButtons.forEach(btn => {
            btn.classList.toggle('active', btn.dataset.lang === lang);
        });

        // Tüm çevrilebilir elementleri güncelle
        document.querySelectorAll('[data-tr], [data-en]').forEach(element => {
            const text = element.dataset[lang];
            if (text) {
                if (element.tagName === 'A') {
                    const span = element.querySelector('span');
                    if (span) {
                        span.textContent = text;
                    } else {
                        element.textContent = text;
                    }
                } else if (element.tagName === 'SPAN') {
                    element.textContent = text;
                } else if (element.tagName === 'CODE') {
                    element.textContent = text;
                    // Prism.js'i yeniden vurgula
                    Prism.highlightElement(element);
                } else {
                    element.textContent = text;
                }
            }
        });

        // HTML lang attribute'unu güncelle
        document.documentElement.lang = lang;
    }

    // Dil butonlarına tıklama olayı ekle
    langButtons.forEach(btn => {
        btn.addEventListener('click', () => {
            changeLang(btn.dataset.lang);
        });
    });

    // Sayfa yüklendiğinde EN dilini seçili yap
    changeLang('en');

    // Smooth scrolling for anchor links
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function (e) {
            e.preventDefault();
            const target = document.querySelector(this.getAttribute('href'));
            if (target) {
                target.scrollIntoView({
                    behavior: 'smooth',
                    block: 'start'
                });
            }
        });
    });

    // Active link highlighting
    const sections = document.querySelectorAll('section');
    const navLinks = document.querySelectorAll('.nav-links a');

    window.addEventListener('scroll', () => {
        let current = '';
        sections.forEach(section => {
            const sectionTop = section.offsetTop;
            const sectionHeight = section.clientHeight;
            if (pageYOffset >= sectionTop - 200) {
                current = section.getAttribute('id');
            }
        });

        navLinks.forEach(link => {
            link.classList.remove('active');
            if (link.getAttribute('href').substring(1) === current) {
                link.classList.add('active');
            }
        });
    });

    // Mobil menü iyileştirmeleri
    const menuToggle = document.createElement('button');
    menuToggle.classList.add('menu-toggle');
    menuToggle.innerHTML = '<i class="fas fa-bars"></i>';
    document.querySelector('.sidebar-header').appendChild(menuToggle);

    const sidebar = document.querySelector('.sidebar');
    const content = document.querySelector('.content');

    menuToggle.addEventListener('click', () => {
        sidebar.classList.toggle('active');
        menuToggle.innerHTML = sidebar.classList.contains('active') 
            ? '<i class="fas fa-times"></i>' 
            : '<i class="fas fa-bars"></i>';
    });

    // Menü dışına tıklandığında menüyü kapat
    document.addEventListener('click', (e) => {
        if (!sidebar.contains(e.target) && !menuToggle.contains(e.target) && sidebar.classList.contains('active')) {
            sidebar.classList.remove('active');
            menuToggle.innerHTML = '<i class="fas fa-bars"></i>';
        }
    });

    // Ekran boyutu değiştiğinde menüyü kontrol et
    window.addEventListener('resize', () => {
        if (window.innerWidth > 768 && sidebar.classList.contains('active')) {
            sidebar.classList.remove('active');
            menuToggle.innerHTML = '<i class="fas fa-bars"></i>';
        }
    });

    // Menü linklerine tıklandığında menüyü kapat
    document.querySelectorAll('.nav-links a').forEach(link => {
        link.addEventListener('click', () => {
            if (window.innerWidth <= 768) {
                sidebar.classList.remove('active');
                menuToggle.innerHTML = '<i class="fas fa-bars"></i>';
            }
        });
    });

    // Code block copy button
    document.querySelectorAll('pre code').forEach((block) => {
        const button = document.createElement('button');
        button.classList.add('copy-button');
        button.innerHTML = '<i class="fas fa-copy"></i>';
        block.parentNode.style.position = 'relative';
        block.parentNode.appendChild(button);

        button.addEventListener('click', () => {
            const code = block.textContent;
            navigator.clipboard.writeText(code).then(() => {
                button.innerHTML = '<i class="fas fa-check"></i>';
                setTimeout(() => {
                    button.innerHTML = '<i class="fas fa-copy"></i>';
                }, 2000);
            });
        });
    });
}); 