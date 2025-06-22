class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-go-0841d2b491c59080b359323264928d20d13fab9b/example-go"
  version "0841d2b491c59080b359323264928d20d13fab9b"

  def install
    bin.install "example-go"

#     system bin/"example-go", "man"
#     man1.install Dir["*.1"]

    generate_completions_from_executable(bin/"example-go", "completion")
  end
end
