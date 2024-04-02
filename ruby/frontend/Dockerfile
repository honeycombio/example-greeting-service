FROM ruby:3.0
RUN gem install bundler
WORKDIR /myapp
COPY Gemfile* /myapp/
RUN bundle install
COPY frontend.ru /myapp
COPY o11y_wrapper.rb /myapp

EXPOSE 7777

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "curl", "--fail", "http://localhost:7777/up" ]
CMD [ "bundle", "exec", "rackup", "frontend.ru", "--server", "puma", "--host", "0.0.0.0"]
