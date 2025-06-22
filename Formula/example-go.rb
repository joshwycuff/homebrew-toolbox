class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-go-0841d2b491c59080b359323264928d20d13fab9b/example-go"
  version "0841d2b491c59080b359323264928d20d13fab9b"

  def install
    bin.install "example-go"
    man1.install Dir["*.1"]
    bash_completion.install "example-go.bash" => "example-go"
    zsh_completion.install "example-go.zsh" => "_example-go"
    fish_completion.install "example-go.fish"
  end
end
