# APKProject
در این پروژه پکیجی به نام 
"webscraper"
درنظر گرفته شده است که دارای دو تابع
ReadWebData,StorageData
تعریف شده است. در تابع اول با استفاده از پکیج 
net/http
دیتا از وبسایت خوانده شد و رشته هر لاین جدا گردید و به صورت آرایه ای از رشته ها ذخیره گردید.
در تابع دوم ابتدا اطلاعات ذخیره شده به صورت یک ساختار مشخص درآمد و سپس با استفاده از پکیج های 
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	esapi "github.com/elastic/go-elasticsearch/v8/esapi"
در ElasticSearch
ذخیره گردید.
در پکیج 
main
با استفاده از پکیج
	"github.com/go-co-op/gocron"
یک زمانبندی برای اجرای تسک نوشته شد و در داخل تسک تابع 
 ReadWebData
 فراخوانی شد .
در نهایت برای ایجاد سرویس لینوکسی دستورات زیر اجرا شد
sudo nano /etc/systemd/system/apk-read.service
[Unit]
Description= Read and Write Elasticsearch

[Service]
Type=simple
User=root
Group=root
ExecStart=/home/shakibacharkhkar/GoLangProjects/APKProject/APKElasticSearch
Restart=always


[Install]
WantedBy=mult-user.target

sudo systemctl daemon-reload
sudo systemctl start apk-read.service
sudu systemctl status apk-read.service


